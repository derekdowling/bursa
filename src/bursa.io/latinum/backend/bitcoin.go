package backend

import (
	"log"

	"bursa.io/latinum/shared/config"
	"bursa.io/latinum/vault"
	"bursa.io/config"
	"github.com/conformal/btcrpcclient"
	"github.com/conformal/btcjson"
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
	vault.GetEncodedAddress()
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
func (self *Latinum) GenerateInto(amt int, encoded_address string) error {
	err := self.client.SetGenerate(true, amt)
	if err != nil {
		log.Fatalf("Couldn't generate bitcoins")
	}

	unspent, err := self.client.ListUnspent()

	// Aggregate unspent transactions until we have more than then requested amount.
	// Who needs ruby? Good old for loops.
	current_amt := 0
	index := 0
	var inputs []btcjson.TransactionInput
	for utxo := range unspent {
		if current_amt > amt {
			break
		}

		current_amt += utxo.amount
		append(btcjson.TransactionInput{Txid: utxo.Txid, Vout: index}, utxo)
		index += 1
	}

	if (curent_amt < amt) {
		return errors.New("Insufficient funds in server wallet")
	}

	// Transaction fee is the difference between in/out.
	tx_fee := 0.00001

	// Calculate change to send back to ourselves.
	change := current_amt - tx_fee - amt

	var inputs []btcjson.TransactionInput
	var amounts map[btcutil.Address]btcutil.Amount

	src_address, err := self.client.GetNewAddress()
	if err != nil {
		return err
	}

	dest_address, err := btcutil.DecodeAddress(encoded_address, config.BTCNet())

	amounts[src_address] = amt
	amounts[dest_address] = change

	unsigned_raw_tx, err := self.client.CreateRawTransaction(inputs, amounts)

	// TODO we may want to return a new error rather than the descended one.
	if err != nil {
		return err
	}

	// This should be around the final step.
	// self.latinum.client.SignRawTransaction(unsigned_raw_tx, 
}
