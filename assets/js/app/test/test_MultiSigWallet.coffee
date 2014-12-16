MultiSigWallet = require 'app/models/MultiSigWallet'
User = require 'app/models/user'

describe 'MultiSigWallet', ->
  describe 'create', ->
    multisig = MultiSigWallet.create(null)
    expect(multisig).to.have.property 'user'
