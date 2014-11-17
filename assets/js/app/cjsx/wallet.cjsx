# @cjsx React.DOM
React = require 'react'
require 'es6-shim' # {... extraProps}

{ Route, Routes, Link, Navigation } = require 'react-router'
{ Button, ButtonGroup } = require 'react-bootstrap'

Balance = require './balance.cjsx'
Hash = require './hash.cjsx'

Wallet = React.createClass
    mixins: [ Navigation ]

    isActive: ->
      unless @props.path?
        return false
      @props.path[-1..]?[0] == @props.address

    isAncestor: ->
      unless @props.path?
        return false
      @props.address in @props.path[...-1]

    highlightClass: ->
      if @isActive() then return 'panel-success' else ''

    isPreview: ->
      unless @props.path?
        return false
      console.log @props
      @props.level == @props.active_level

    classes: ->
      [
        if @isActive() then "active"
        if @isAncestor() then "ancestor"
        if @isPreview() then "preview"
      ].join(" ")

    render: ->
      wallet_link = @makeHref "/wallets/wallet/#{@props.address}"

      <div className="wallet col-md-3 #{@classes()}">
        <div className="panel panel-default #{@highlightClass()}">
          <div className="panel-heading">
            <Link to="/wallets/#{@props.address}"><h4>{@props.label}</h4></Link>
            <Balance balance={@props.balance} />
          </div>
          <div className="panel-body">
            <Hash hash={@props.address} />
          </div>
          <div className="panel-footer">
            <ButtonGroup className="tools">
              <Button><i className="fa fa-plus-circle"/> New</Button>
              <Button href={wallet_link}><i className="fa fa-plus-circle"/> View</Button>
            </ButtonGroup>
          </div>
        </div>
        <div className="ancestor-arrow">
          <i className="fa fa-3x fa-chevron-down"/>
        </div>
      </div>

module.exports = Wallet
