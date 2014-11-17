# @cjsx React.DOM
React = require 'react'

{ Route, Routes, Link, ActiveState } = require 'react-router'

module.exports = NavLink = React.createClass
  mixins: [ ActiveState ]

  render: ->
    isActive = @isActive(@props.to, @props.params, @props.query)
    className = if isActive  then 'active' else ''
    <li className={className}><Link {... @props}/></li>
