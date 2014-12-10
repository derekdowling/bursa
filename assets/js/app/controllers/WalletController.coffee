expect         = require('chai').expect
Promise        = require 'bluebird'
{ Wallet
  HDNode }     = require 'bitcoinjs-lib'

request        = require 'app/lib/request'
config         = require 'app/config'
Controller     = require 'app/controllers/Controller'
MultiSigWallet = require 'app/models/MultiSigWallet'
User           = require 'app/models/User'

# Strictly a restful controller. Does not do frontend things.
class WalletController extends Controller
  # Automatically configures our local db if present.
  @db_schema =
    wallets: 'identifier, owner'

  @resource = 'wallets'

  # Saves the external wallet provided to bursa.
  saveExternal: (owner_id, wallet, pubkeys) ->
    expect(wallet).to.be.instanceOf HDNode
    expect(owner_id).to.be.a 'string'
    expect(pubkeys).to.be.instanceOf Array

    @request('put').send(wallet: wallet.toHex(), owner: owner_id, pubkeys: pubkeys)

  # Saves the local wallet for the user.
  saveLocal: (owner_id, wallet, pubkeys) ->
    expect(wallet).to.be.instanceOf HDNode
    expect(owner_id).to.be.a 'string'
    expect(pubkeys).to.be.instanceOf Array

    @table.add(id: wallet.getIdentifier(), data: wallet.toHex(), owner: owner_id)

  getLocal: (owner_id, id) ->
   @table
      .where("owner")
      .equals(owner_id)
      .and("id")
      .equals(id)

  # Creates a multi-sig wallet and saves the local as well
  # as external keys. The cold storage and local key's are returned
  # for further processing.
  create: (owner, parent_wallet=null) ->
    expect(owner).to.be.instanceOf User

    Promise.resolve().then =>
      MultiSigWallet.create(parent_wallet)
    .then (result) =>
      { user, server, cold, pubkeys } = result

      Promise.all([
        @saveLocal owner.id, user, pubkeys
        @saveExternal owner.id, server, pubkeys
      ])
      .then -> { user, cold, pubkeys }

module.exports = WalletController
