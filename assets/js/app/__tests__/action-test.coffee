jest.dontMock('../actions/WalletActions')
{ WalletAction, WalletCreateAction } = require '../actions/WalletActions'

describe 'Action Tests', ->
  it 'Creates an action fluently', ->
    expect(WalletCreateAction.build(hash:5)).toEqual(jasmine.any(WalletCreateAction))
