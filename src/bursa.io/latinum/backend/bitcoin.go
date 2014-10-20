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
