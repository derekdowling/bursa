unless localStorage?
  localStorage = require 'node-localstorage'

request        = require 'app/lib/request'
config         = require 'app/config'

# Intended to bridge the gap between our Stores, which model application state,
# and various external services - which help persist or read data.
class Controller
  @api_version    = 1
  @resource       = null
  @host           = config.api.host

  constructor: ({@db, @request}={}) ->
    # Base HTTP request driver.
    unless @request?
      @request = (method, url) ->
        urlparts = [
          @constructor.api_version
          @constructor.resource
          url
        ].filter (n) -> typeof n != 'undefined'

        (new request.Request(
          method,
          "#{@constructor.host}/#{urlparts.join('/')}"
        )).type('json')

    # Local indexdb storage
    if @constructor.db_schema?
      @db ?= new Dexie(@constructor.resource)
      @db.version(@constructor.db_version).stores(@constructor.db_schema)
      @db.open()

    @table = @db[@constructor.resource]

module.exports = Controller
