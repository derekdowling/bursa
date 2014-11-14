# @cjsx React.DOM
React = require 'react'

Nav = React.createClass
    render: ->
      <div className="nav-wrapper">
        <ul id="nav" className="nav" data-slim-scroll data-collapse-nav data-highlight-active>
          <li><a href="#/wallets"><i className="fa fa-bitcoin"></i><span>Wallets</span></a></li>
          <li><a href="#/send"><i className="fa fa-send"></i><span>Send</span></a></li>
          <li><a href="#/receive"><i className="fa fa-reply"></i><span>Receive</span></a></li>
          <li><a href="#/delegate"><i className="fa fa-share-alt"></i><span>Delegate</span></a></li>
        </ul>
      </div>

React.render <Nav />, document.getElementById 'nav-container'

module.exports = Nav
