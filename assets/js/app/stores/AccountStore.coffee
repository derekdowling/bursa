AppDispatcher    = require '../dispatchers/Dispatcher'
Store            = require './Store'

class AccountStore extends Store
  initialize: ->

  # @see WalletViewCreateAction
  onViewAuthorizeAction: (action) ->

# We expose a singleton instead of the class itself. Potentially, if we want to
# migrate to full-di at some point, this will make our task easier.
# Also makes subclassing harder.
accountStore = new AccountStore()

AppDispatcher.registerStore accountStore

module.exports = walletStore
