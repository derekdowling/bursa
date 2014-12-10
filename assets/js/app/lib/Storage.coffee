unless localStorage?
  localStorage = require 'node-localstorage'

class Storage
  storeWallet: (wallet) ->

