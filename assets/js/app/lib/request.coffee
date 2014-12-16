# Augments superagent with append for url appends.
request = require 'superagent'
Uri = require 'URIjs'

request.Request.prototype.append = (parts...)->
  @url = (new Uri(@url)).segment(parts).readable()
  @

# As of Dec 2014 these available on the proxy object but not the Request
# class, so defining them should be okay.
Object.defineProperty request.Request.prototype, 'put',
  get: ->
    @method = 'put'
    @

Object.defineProperty request.Request.prototype, 'get',
  get: ->
    @method = 'get'
    @

Object.defineProperty request.Request.prototype, 'post',
  get: ->
    @method = 'post'
    @

module.exports = request
