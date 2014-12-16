Dispatcher = require 'app/dispatchers/Dispatcher'

class Action
  # Fluently delegates to a constructor.
  @build: ->
    new @(arguments...)

  dispatch: ->
    Dispatcher.dispatch @

  name: ->
    "#{@constructor.name}"

class ViewAction extends Action

class ServerAction extends Action

module.exports = { Action, ViewAction, ServerAction }
