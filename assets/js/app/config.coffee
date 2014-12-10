{ networks } = require 'bitcoinjs-lib'

module.exports = {
  bitcoin:
    network: networks.testnet
  api:
    host: 'https://dev.bursa.io'
}
