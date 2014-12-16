expect = require('chai').expect
{ ViewAction } = require 'app/actions/Actions'

class WalletViewAction extends ViewAction

class WalletViewCreateAction extends WalletViewAction
  # The address of the parent wallet.
  constructor: (@parentAddress) ->
    expect(@parentAddress).to.be.a('string')

class WalletViewDestroyAction extends WalletViewAction
  constructor: (@wallet) ->

module.exports = { WalletViewAction, WalletViewCreateAction, WalletViewDestroyAction }
