// Responsible for secure private key management. It currently also has some
// convenience methods for generating public keys and addresses which may not
// belong here - the vault shouldn't need to be interacted with as frequently as
// public keys in general. So putting that functionality in promiximity to
// private key related operations may just tempt us to undermine the security
// oriented intentions of this package.
package vault

import (
	"bursa.io/latinum/vault/store"
	"bursa.io/latinum/backend/client"
	"bursa.io/latinum/shared/config"

	"github.com/conformal/btcutil/hdkeychain"
	"github.com/conformal/btcwire"
	"github.com/conformal/btcutil"
	"log"
)

func GetPublicKey(username string) string {
	return "Stub"
}

// Sign signs a transaction with the private key specified. Once signed,
// a transaction may be ready for processing unless additional signatures are
// required.
//
// TODO passing the latinum client wrapper is just a crappy side effect of wanting
// to make this work for now. This should be isolated from the network and needing
// an rpc call to bitcoind in order to sign a transaction input is extremely MEH.
func SignByUserId(tx *btcwire.MsgTx, user_id int64) (*btcwire.MsgTx, error) {
	// TODO error checking
	encoded_key, _ := store.Retrieve(user_id)

	return SignWithEncodedExtendedKey(tx, encoded_key)
}

func signWithWIFKey(tx *btcwire.MsgTx, wif_key *btcutil.WIF) (*btcwire.MsgTx, error) {
	// TODO I really wanted to just do this in-line.
	var private_keys []string
	private_keys = append(private_keys, wif_key.String())

	// NOTE The ok parameter "indicates whether the received value was sent on the
	// channel (true) or is a zero value returned because the channel is closed and
	// empty (false)."
	response, _, err := client.Get().SignRawTransaction3(tx, nil, private_keys)
	return response, err
}

func SignWithEncodedWIFKey(tx *btcwire.MsgTx, encoded_key string) (*btcwire.MsgTx, error) {
	wif_key, err := btcutil.DecodeWIF(encoded_key)
	if err != nil {
		log.Fatalf("Couldn't decode encoded wif key", err)
	}

	return signWithWIFKey(tx, wif_key)
}

func SignWithEncodedExtendedKey(tx *btcwire.MsgTx, encoded_key string) (*btcwire.MsgTx, error) {
	// This beautiful sequence converts the encoded private key into a
	// Wallet Import Format (WIF) private key that the rpc client can use.
	// The API, as you can see, is garbage.
	extended_key, err := hdkeychain.NewKeyFromString(encoded_key)
	if err != nil {
		log.Fatalf("Couldn't create extended key", err)
	}

	private_key, err := extended_key.ECPrivKey()
	if err != nil {
		log.Fatalf("Couldn't create private key", err)
	}

	wif_key, err := btcutil.NewWIF(
		private_key,
		config.BTCNet(),
		true,
	)

	if err != nil {
		log.Fatalf("Couldn't convert key to WIF", err)
	}

	return signWithWIFKey(tx, wif_key)
}

// Generate a new private key for a given user.
// TODO error propagation.
func NewMasterForUser(user_id int64) (string, error) {
	key, _ := NewMaster()
	store.Store(user_id, key)
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
func GetEncodedAddressForUser(user_id int64) (string) {
	encoded_key, err := store.Retrieve(user_id)

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
func GetEncodedAddress(encoded_base_58_key string) string {
	key, err := hdkeychain.NewKeyFromString(encoded_base_58_key)
	// TODO look through all my Fatalf's and handle them gracefully.
	if err != nil {
		log.Fatalf("Failed to decode key", err)
	}

	address, err := key.Address(config.BTCNet())
	if err != nil {
		log.Fatalf("Failed to decode key", err)
	}

	// It appears a public key (which is what an address *kinda* is) can take on
	// many forms.
	// TODO How are addresses different from public keys, or what makes a public
	// key into a valid address?
	return address.String()
}
