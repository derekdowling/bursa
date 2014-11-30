Dispatcher = require '../dispatchers/Dispatcher'
expect = require('chai').expect

class WalletAction
  # Fluently delegates to a constructor.
  @build: ->
    new @(arguments...)

  dispatch: ->
    Dispatcher.handleViewAction @

class WalletCreateAction extends WalletAction
  constructor: (@parentWallet) ->
    expect(@parentWallet).to.have.property('hash')

class WalletDestroyAction extends WalletAction
  constructor: (@wallet) ->

  dispatch: ->
    Dispatcher.handleViewAction @

module.exports = { WalletAction, WalletCreateAction, WalletDestroyAction }
