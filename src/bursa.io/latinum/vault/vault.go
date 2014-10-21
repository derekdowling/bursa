// Responsible for secure private key management. It currently also has some
// convenience methods for generating public keys and addresses which may not
// belong here - the vault shouldn't need to be interacted with as frequently as
// public keys in general. So putting that functionality in promiximity to
// private key related operations may just tempt us to undermine the security
// oriented intentions of this package.
package vault

import (
	"bursa.io/latinum/vault/store"
	"github.com/conformal/btcnet"
	"github.com/conformal/btcutil/hdkeychain"
	"log"
)

// Network parameters are necessary during the derivation of public keys
// from extended keys. They vary between mainnet and regtest.
// TOLEARN Which parameters exactly are used during the derivation? How do they
// affect things?
//
// TODO This flag being wrong in production could represent a signficant risk.
// Would money be lost if we tried to sign transfers using bunk keys? Or would
// they simply "bounce"?
var btc_network *btcnet.Params

func init() {
	// TODO PROD
	// btc_network = &btcnet.MainNetParams
	btc_network = &btcnet.RegressionNetParams
}

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

// Generate a new private key for a given user.
// TODO error propagation.
func NewMasterForUser(user_id int64) (string, error) {
	key = NewMaster()
	store.Store(user_id, key.String())
	return key, nil
}

// Generate a new private key.
// TODO error propagation.
func NewMaster() (string, error) {
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

	return key.String(), nil
}

// Currently does not take advantage of HD functionality. We simply return
// a public key associated with the master private key and convert it to a base58
// encoded bitcoin address.
//
// This is potentially a huge security risk, from BIP32:
// Knowledge of the extended public key plus any non-hardened private key
// descending from it is equivalent to knowing the extended private key (and thus
// every private and public key descending from it).
//
// The use case for non-hardened might be auditing it seems? Share a public key
// at a given depth in the organization with an auditor and they can see all
// transactions made to any descended public key, but cannot spend your money?
func GetEncodedAddressForUser(user_id int64) (string, error) {
	encoded_key, err := Retrieve(user_id)

	// TODO this is harsh. It will happen if the user simply doesn't have a key.
	// We don't wait it killing our entire daemon in that case.
	// Look into better error handling.
	if err != nil {
		log.Fatalf("Failed to retrieve key", err)
	}

	return GetEncodedAddress(encoded_key)
}

// Returns an encoded public address hash (usable with P2PKH) for a given encoded
// private key.
func GetEncodedAddress(encoded_base_58_key string) (string, error) {
	key, err := hdkeychain.NewKeyFromString(encoded_base_58_key)
	// TODO look through all my Fatalf's and handle them gracefully.
	if err != nil {
		log.Fatalf("Failed to decode key", err)
	}

	address, err := key.Address(btc_network)
	if err != nil {
		log.Fatalf("Failed to decode key", err)
	}

	// It appears a public key (which is what an address *kinda* is) can take on
	// many forms.
	// TODO How are addresses different from public keys, or what makes a public
	// key into a valid address?
	return address.Encode()
}
