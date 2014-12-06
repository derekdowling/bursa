Dispatcher = require '../dispatchers/Dispatcher'
expect = require('chai').expect

class WalletAction
  # Fluently delegates to a constructor.
  @build: ->
    new @(arguments...)

  dispatch: ->
    Dispatcher.dispatch @

class WalletCreateAction extends WalletAction
  # The address of the parent wallet.
  constructor: (@parentAddress) ->
    expect(@parentAddress).to.be.a('string')

class WalletDestroyAction extends WalletAction
  constructor: (@wallet) ->

  dispatch: ->
    Dispatcher.handleViewAction @

module.exports = { WalletAction, WalletCreateAction, WalletDestroyAction }
