WalletController = require 'app/controllers/WalletController'
MultiSigWallet = require 'app/models/MultiSigWallet'
User = require 'app/models/User'

Dexie = require 'dexie'
{ indexedDB } = require 'indexeddb-js'
sqlite  = require 'sqlite3'
Dexie.dependencies.indexedDB = new indexedDB('sqlite3', new sqlite.Database(':memory'))

{ HDNode, Wallet, networks } = require "bitcoinjs-lib"


describe 'WalletController', ->
  request = null
  db = null

  beforeEach ->
    request = ->
      end: sinon.stub().onCall(0).resolves({ok: true})
      send: sinon.stub().returnsThis()

    db =
      open: sinon.stub().returnsThis()
      version: sinon.stub().returnsThis()
      stores: sinon.stub().returnsThis()
      wallets:
        add: sinon.stub().returnsThis()


  describe 'API', ->
    describe 'Save', ->
      spy = null
      controller = null
      user = null

      beforeEach ->
        user = new User "test1"
        controller = new WalletController request: request, db: db

      it 'should have a table', ->
        expect(controller.table).to.not.eql null
        expect(controller.table).to.not.eql undefined

      it 'should save externally', ->
        expect(controller.create(user)).to.eventually.to.fulfilled

      it 'should save locally', ->
        expect(controller.create(user)).eventually.to.be.fulfilled

      it 'should return user keys private keys the public keys', ->
        expect(controller.create(user)).to.eventually.be.fulfilled.then (result) ->
          expect(result).to.have.property 'user'
          expect(result).to.have.property 'cold'
          expect(result).to.have.property 'pubkeys'
          expect(result).to.not.have.property 'server'
          expect(result.pubkeys).to.have.length 3

  describe 'Integration', ->
    describe 'Save', ->
      it 'should save locally', ->
        owner = new User "test1"
        controller = new WalletController()
        { user, server, cold, pubkeys } = MultiSigWallet.create()
        controller.saveLocal owner.id, user, pubkeys
