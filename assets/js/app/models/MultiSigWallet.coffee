expect = require('chai').expect
{ Wallet, networks } = require 'bitcoinjs-lib'
config = require 'app/config'

# A 2 of 3 multisig wallet.
class MultiSigWallet
  constructor: (@pub_keys) ->
    expect(@pub_keys).to.have.length(3)

  # Build a multi sig wallet. Provide 2 keys to the user, one of which is saved
  # in local storage, the other which the user is instructed to place in cold
  # storage or onto a hardware key.
  #
  # The remaining wallet is sent to Bursa for us to hold on to.
  #
  # In order to compromise a wallet, the user's wallet and Bursa would need to be
  # compromised.
  #
  # In the event that Bursa dies, the user can still release their funds via
  # their cold storage key.
  @create: (parent_wallet=null) ->
    # The root wallet which derives our children. Only intended for testing.
    parent_wallet ?= new Wallet(null, config.bitcoin.network)
    expect(parent_wallet).to.be.an.instanceOf Wallet

    # Return the derived wallets. They should not persist for long.
    # The user's primary wallet.
    user = parent_wallet.getExternalAccount().deriveHardened(0)
    # A second wallet we store on the user's behalf.
    server = parent_wallet.getExternalAccount().deriveHardened(1)
    # A third wallet that allows the user to override us.
    cold = parent_wallet.getExternalAccount().deriveHardened(2)
    # The public keys are distributed to all wallets so they
    # can identify transfers involving the respective private keys.
    pubkeys = [
      user.pubKey.toHex()
      server.pubKey.toHex()
      cold.pubKey.toHex()
    ]

    {user, server, cold, pubkeys }

module.exports = MultiSigWallet
