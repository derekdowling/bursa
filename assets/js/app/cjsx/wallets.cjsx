# @cjsx React.DOM
React = require 'react'
require 'es6-shim' # {... extraProps}

{ Route, Routes, Link } = require 'react-router'

Wallet = require './wallet.cjsx'
WalletStore = require '../stores/WalletStore'


# wallets is a hash of wallets keyed by their address.
# path is a slash delimited string.
#
# Returns sibling wallets at each level of the hierarchy.
#                     a
#            1                   2
#      d     e     f       h     i     j
#     4 5   6 7   8 9     10
#
# E.g. a/1/d/4 will yield:
#
#   [ [a], [1, 2], [d, e, f], [4, 5] ]
collectLevels = (wallets, path = null) ->
  unless path?
    return [(wallet for key, wallet of wallets)]
  unless wallets?
    return []

  levels = for address, level in path
    level_wallets = (wallet for key, wallet of wallets)
    wallets = wallets[address].wallets
    level_wallets

  # Provide the next level of children
  if Object.keys(wallets).length
    levels[level] = (wallet for key, wallet of wallets)
  levels

pathToWallet = (wallets, address) ->
  unless address?
    return null

  dfs = (wallet, stack) ->
    stack.push wallet.address
    if wallet.address == address
      return stack
    else
      for key, wallet of wallet.wallets
        result = dfs(wallet, stack[..])
        if result then return result

    return null

  for key, wallet of wallets
    if found = dfs(wallet, [])
      return found

  return null

Wallets = React.createClass
    getInitialState: ->
      wallets:
        "1KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
          label: "Bursa.io"
          address: "1KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
          balance: 50.0
          wallets:
            "2KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
              label: "Capital"
              address: "2KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
              balance: 20.0
              wallets:
                "5KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
                  label: "Infrastructure"
                  address: "5KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
                  balance: 14.65
                  wallets: {}
                "6KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
                  label: "Equipment"
                  address: "6KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
                  balance: 5.35
                  wallets: {}
            "3KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
              label: "Marketing"
              address: "3KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
              balance: 21.29
              wallets: {}
            "4KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c":
              label: "Payroll"
              address: "4KNp2RrFvtRLh7FX6qAwYzqN6d1bmM849c"
              balance: 8.71
              wallets: {}

    # So this is a bit  ugly. These calcuations could be done before hand
    # wallets shouldn't receive random loop indices, state, props blah.
    deriveState: (nextProps, nextState) ->
      path = pathToWallet nextState.wallets, nextProps.params.address
      levels = collectLevels nextState.wallets, path
      active_level = path?.length or 0

      wallets: nextState.wallets
      path: path
      levels: levels
      active_level: active_level

    componentWillReceiveProps: (nextProps) ->
      @setState(@deriveState(nextProps, @state))

    componentWillMount: ->
      @setState(@deriveState(@props, @state))
      WalletStore.addChangeListener ->

    componentWillUnmount: ->
      WalletStore.removeChangeListener ->

    renderLevel: (wallets, depth) ->
      <div className="row wallet-level level-#{depth}">
        { for wallet in wallets
          <Wallet {... wallet}
            key={wallet.address}
            path={@state.path}
            level={depth}
            active_level={@state.active_level}
            />
        }
      </div>

    renderLevels: ->
      for level, i in @state.levels
        @renderLevel level, i

    render: ->
      wallets = @renderLevels()
      <div className="wallets">
        <div className="row">
          <div className="col-sm-12">
              <h2>Wallets</h2>
          </div>
        </div>
        {wallets}
      </div>

module.exports = Wallets
