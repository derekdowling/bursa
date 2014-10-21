package backend

import (
	"log"

	"bursa.io/config"
	"github.com/conformal/btcrpcclient"
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
func (self *Latinum) GenerateInto(amt int, address string) error {
	err := self.client.SetGenerate(true, amt)
	if err != nil {
		log.Fatalf("Couldn't generate bitcoins")
	}
}
