package backend

import (
	"log"
	"fmt"
	"errors"

	"bursa.io/latinum/backend/client"
	shared_config "bursa.io/latinum/shared/config"
	"bursa.io/latinum/vault"

	"bursa.io/config"
	"github.com/conformal/btcrpcclient"
	"github.com/conformal/btcjson"
	"github.com/conformal/btcutil"
	"github.com/spf13/viper"
)

func init() {
	// I feel like this piece should be part of some kind of applicatio kernel it
	// would be nice to not have to load it while also avoiding becoming to dependent
	// on said application kernel to establish a working state for the application
	// - which tends to lead to annoying-to-test code.
	config.LoadConfig()
}

// Not quite sure what we want here yet
type TransferResponse struct {
	msg  string
	code int
}

type Transfer struct {
}

type CurrencyBackend interface {
	ExecuteTransfer(transfer *Transfer) *TransferResponse
}

type Latinum struct {
	client *btcrpcclient.Client
}

func NewLatinum() *Latinum {
	client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
		// There appears to be a bug with nested booleans and viper.GetBool.
		// HttpPostMode: viper.GetBool("bitcoin.HttpPostMode"),
		HttpPostMode: true,
		DisableTLS:   true,
		Host:         viper.GetString("bitcoin.Host"),
		User:         viper.GetString("bitcoin.User"),
		Pass:         viper.GetString("bitcoin.Pass"),
	}, nil)

	// I usually like to connect explicitly in a Connect call.
	// This kind of unexpected bailing is pretty side-effect-y IMO.
	if err != nil{
		log.Fatalf("Latinum could not create a new Bitcoin client.", err)
	}

	return &Latinum{client: client}
}

// Sends amt satoshis from `from_address` to `to_address`. Automatically
// calculates change output. Uses a greedy strategy of using your smallest UTXO's
// first.
func (self *Latinum) Send(from_address string, to_adddress string, amt int) {
}

// Generates bitcoins and gives them to an address. Used only during testing
// in regtest mode where setgenerate is actually functional.
//
// When we generate our initial bitcoins for test purposes, they're implicitly
// associated with bitcoind's default wallet. Generally, all of Bursa's other
// operations take the burden of wallet management out of bitcoind.
//
// Amt shoult be set to > 101 in order to ensure we've confirmed some blocks.
// to make the newly minted bitcoins spendable. Yes - we're "mining" here. The
// 100 confirmations thing is to iron out blockchain forks.
func GenerateInto(amt float64, encoded_address string) error {
	// err := client.Get().SetGenerate(true, 100)
	// if err != nil {
	//   log.Fatalf("Couldn't generate bitcoins", err)
	// }

	unspent, err := client.Get().ListUnspent()

	// Aggregate unspent transactions until we have more than then requested amount.
	// Who needs ruby? Good old for loops.
	current_amt := 0.0
	var inputs []btcjson.TransactionInput
	var amounts = make(map[btcutil.Address]btcutil.Amount)

	for _, utxo := range unspent {
		fmt.Println(utxo.Address)
		if utxo.Address != "mnFWmpz1xYi1SeR8KkSyPD5es4TkY2LTto" {
			continue
		}

		if current_amt > amt {
			break
		}

		inputs = append(inputs, btcjson.TransactionInput{Txid: utxo.TxId, Vout: utxo.Vout})
		current_amt += utxo.Amount
	}

	if (current_amt < amt) {
		return errors.New("Insufficient funds in server wallet")
	}

	fmt.Println("current amt:", current_amt)

	// Transaction fee is the difference between in/out.
	tx_fee := 0.001

	// Calculate change to send back to ourselves.
	change := current_amt - tx_fee - amt

	src_address, err := client.Get().GetNewAddress()
	if err != nil {
		return err
	}

	fmt.Println("change", change)
	fmt.Println("amt", amt)
	fmt.Println("tx_fee", tx_fee)
	fmt.Println("new address", src_address)
	fmt.Println("encoded address", encoded_address)

	dest_address, err := btcutil.DecodeAddress(encoded_address, shared_config.BTCNet())
	if err != nil {
		log.Print("Couldn't decode destination address", err)
		return err
	}

	amounts[src_address], _ = btcutil.NewAmount(amt)
	amounts[dest_address], _ = btcutil.NewAmount(change)

	unsigned_raw_tx, err := client.Get().CreateRawTransaction(inputs, amounts)

	fmt.Println(inputs)

	// TODO we may want to return a new error rather than the descended one because
	// it ends up leaking the underlying abstraction details to our caller.
	if err != nil {
		log.Print("Couldn't generate unsigned raw transaction", err)
		return err
	}

	// WIF is the format returned by bitcoin-cli dumpprivkey
	signed, err := vault.SignWithEncodedWIFKey(unsigned_raw_tx, "cP1UYV2etniLmJAoEj9bE88P7PprMHGBCwUvSo6iqiEk3TVFCXWC")
	if err != nil {
		log.Print("Couldn't sign the damn thing.", err)
		return err
	}

	fmt.Println("signed")
	sha_hash, err := client.Get().SendRawTransaction(signed, false)

	fmt.Println(sha_hash)

	return err
}
