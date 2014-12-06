{ Wallet } = require 'bitcoinjs-lib'
expect = require('chai').expect

# This class wraps a wallet with additional descriptive details that are bursa
# specific. The required `wallet` parameter should be a bitcoinjs wallet object.
class BursaWallet
  constructor: (@wallet, { @label }) ->
    expect(@wallet).to.be.a(bitcjoinjs.Wallet)

