{EventEmitter } = require 'events'

class Store extends EventEmitter
  constructor: ->
    @changeEventName = "change.#{@constructor.name}"

  addChangeListener: (callback) ->
    @addListener @changeEventName, callback

  removeChangeListener: (callback) ->
    @removeListener @changeEventName, callback

  emitChange: ->
    @emit @changeEventName
