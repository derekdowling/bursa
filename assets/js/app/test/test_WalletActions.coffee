{ WalletViewAction, WalletViewCreateAction } = require 'app/actions/WalletActions'

describe 'Action Tests', ->
  it 'Creates an action fluently', ->
    expect(WalletViewCreateAction.build("5")).to.be.instanceOf(WalletViewCreateAction)
