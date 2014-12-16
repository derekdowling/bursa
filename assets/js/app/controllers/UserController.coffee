expect         = require('chai').expect
Promise        = require 'bluebird'
{ Wallet
  HDNode }     = require 'bitcoinjs-lib'

request        = require 'app/lib/request'
config         = require 'app/config'
Controller     = require 'app/controllers/Controller'
MultiSigWallet = require 'app/models/MultiSigWallet'
User           = require 'app/models/User'

class UserController extends Controller
  @resource = 'user'

  save: (user) ->
    expect(user).to.be.instanceOf User
    @request('post').send(user.serialize())

  authenticate: (username, password) ->
    expect(user).to.be.a 'string'
    expect(password).to.be.a 'string'

    @request('post', 'authenticate').send(
      username: username
      password: password
    )
