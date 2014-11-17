React = require 'react'
Router = require 'react-router'

{ Route, Routes, Link } = Router

Viewport = require './viewport.cjsx'
Wallets  = require './cjsx/wallets.cjsx'
WalletDetail  = require './cjsx/wallet_detail.cjsx'
Transfers  = require './cjsx/transfers.cjsx'

# Generate a route that matches nested variables.
dynamic_level_route = (name, depth) ->
  (for i in [1..depth]
    "?:#{name+i}?"
  ).join "/"

address_levels = dynamic_level_route "address", 10

module.exports = routes = (
  <Route handler={Viewport}>
    <Route path="/wallets/?:address?" handler={Wallets}/>
    <Route path="/wallets/wallet/:address" handler={WalletDetail}/>
    <Route path="/transfers" handler={Transfers}/>
    <Route path="/delegate"/>
  </Route>
)


React.render <Routes children={routes}/>, document.getElementById 'body'
