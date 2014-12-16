Promise = require 'bluebird'
Dispatcher = require('flux').Dispatcher

class AppDispatcher extends Dispatcher
  # A wrapper that saves us the boilerplate of assigning the dispatch token as
  # well as delegating to the stores onAction method rather than providing a
  # callback function. Essentially to enshrine a few conventions.
  registerStore: (store) ->
    store.token = @register store.onAction.bind(store)

# Expose a singleton instance.
appDispatcher = new AppDispatcher()

module.exports = appDispatcher
