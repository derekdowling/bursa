jest.dontMock('app/actions/WalletActions.coffee')
{ WalletViewAction, WalletViewCreateAction } = require 'app/actions/WalletActions.coffee'

describe 'Action Tests', ->
  it 'Creates an action fluently', ->
    expect(WalletViewCreateAction.build("5")).toEqual(jasmine.any(WalletViewCreateAction))
