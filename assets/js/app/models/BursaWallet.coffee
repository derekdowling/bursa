expect = require('chai').expect

# This class wraps a wallet with additional descriptive details that are bursa
# specific. The required `wallet` parameter should be a bitcoinjs wallet object.
class BursaWallet
  constructor: (@identifier, { @label }) ->
    expect(@identifier).to.be.a 'string'

module.exports = BursaWallet
