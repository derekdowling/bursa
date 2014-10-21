package vault

import (
	"bursa.io/latinum/vault/store"
	"github.com/conformal/btcutil/hdkeychain"
	"log"
)

func GetPublicKey(username string) string {
	return "Stub"
}

// Sign signs a transaction with the private key specified. Once signed,
// a transaction may be ready for processing unless additional signatures are
// required.
//
// TODO `tx interface` is a placeholder for an actual transaction type - which might
// be provided by one of conformal's libs.
func Sign(tx interface{}) string {
	return "Stub"
}

func NewMaster(user_id int64) string {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		log.Fatalf("Could not generate a new seed.", err)
	}

	// "There is an extremely small chance (< 1 in 2^127) that the seed will derive
	// an unusuable key" - according to the hdkeychain docs. We could do something
	// like retry but umm 1 / 2^127 is pretty good.
	key, err := hdkeychain.NewMaster(seed)
	if err != nil {
		log.Fatalf("Could not generate a master with the given seed", err)
	}

	// Store the key for this user.
	store.Store(user_id, key.String())
	return "key"
}
