class User
  constructor: (@id) ->

  serialize: ->
    id: @id

module.exports = User
