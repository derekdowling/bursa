# @cjsx React.DOM
React = require 'react'

Nav = React.createClass
    render: ->
      <div className="nav-wrapper">
        <ul id="nav" className="nav" data-slim-scroll data-collapse-nav data-highlight-active>
          <li><a href="#/dashboard"><i className="fa fa-dashboard"></i><span>Dashboard</span></a></li>
          <li>
              <a href="#/ui"><i className="fa fa-magic"></i><span>UI Kit</span></a>
              <ul>
                  <li><a href="#/ui/buttons"><i className="fa fa-caret-right"></i><span data-i18n="Buttons"></span></a></li>
                  <li><a href="#/ui/typography"><i className="fa fa-caret-right"></i><span data-i18n="Typography"></span></a></li>
                  <li><a href="#/ui/widgets"><i className="fa fa-caret-right"></i><span data-i18n="Widgets"></span> <span className="badge badge-success">13</span></a></li>
                  <li><a href="#/ui/grids"><i className="fa fa-caret-right"></i><span data-i18n="Grids"></span></a></li>
                  <li><a href="#/ui/icons"><i className="fa fa-caret-right"></i><span data-i18n="Icons"></span></a></li>
                  <li><a href="#/ui/components"><i className="fa fa-caret-right"></i><span data-i18n="Components"></span> <span className="badge badge-danger">18</span></a></li>
                  <li><a href="#/ui/timeline"><i className="fa fa-caret-right"></i><span data-i18n="Timeline"></span></a></li>
                  <li><a href="#/ui/nested-lists"><i className="fa fa-caret-right"></i><span data-i18n="Nested Lists"></span></a></li>
                  <li><a href="#/ui/pricing-tables"><i className="fa fa-caret-right"></i><span data-i18n="Pricing Tables"></span></a></li>
                  <li><a href="#/ui/maps"><i className="fa fa-caret-right"></i><span data-i18n="Maps"></span>  <span className="badge badge-warning">2</span></a></li>
              </ul>
          </li>
        </ul>
      </div>

React.render <Nav />, document.getElementById 'nav-container'

module.exports = Nav
