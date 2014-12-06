jest.dontMock '../stores/WalletStore'
jest.dontMock 'bitcoinjs-lib'
jest.dontMock '../models/BursaWallet'
jest.autoMockOff()

{ Wallet, networks } = require 'bitcoinjs-lib'
WalletStore = require '../stores/WalletStore'
BursaWallet = require '../models/BursaWallet'

describe 'Wallet Tests', ->
  it 'Finds the root wallet', ->
    expect(WalletStore.findWallet("1KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c").label)
      .toEqual("Bursa.io")
  it 'Finds a level 2 wallet', ->
    expect(WalletStore.findWallet("2KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c").label)
      .toEqual("Capital")
  it 'Finds a level 2 sibling wallet', ->
    expect(WalletStore.findWallet("3KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c").label)
      .toEqual("Marketing")
  it 'Finds a level 3 wallet', ->
    expect(WalletStore.findWallet("6KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c").label)
      .toEqual("Equipment")

describe 'WalletDescriptor', ->
  beforeEach ->
    @wallet = new Wallet null, networks.testnet
    @bursaWallet = new BursaWallet @wallet, { label: "awesome" }

  it 'Should wrap a bitcjoinjs wallet', ->
    expect(@bursaWallet.wallet).toBe jasmine.any(Wallet)
    expect(@bursaWallet.label).toEql "awesome"

  it 'Should derive an extended private key', ->
    # 1. Create a wallet using a known seed.
    # 2. Derive a private key from it.
    # 3. Verify that the derived extended private key has
    #    the expected key.
    #
    # FAQ.
    # 1. How do we generate keys consistently, without clashing with
    #    other bitcjoin wallets?
    # A - That's what the seed is for.

  it 'Should derive a neutered private key'
  it 'Should derive an extended public key'
  it 'Should derive a neutered public key'
