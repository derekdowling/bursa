/** @jsx React.DOM */

var Grid = require('./grid.jsx');
var Row = require('./row.jsx');
var Col = require('./col.jsx').Col;

var PricingTableContainer = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <Grid className='container-sm-height'>
        <Row className='row-sm-height'>
          {this.props.children}
        </Row>
      </Grid>
    );
  }
});

var PricingTable = React.createClass({
  propTypes: {
    preferred: React.PropTypes.bool
  },
  render: function() {
    var classes = React.addons.classSet({
      'col-sm-height': true,
      'pricing-table': true,
      'preferred': this.props.preferred
    });
    return this.transferPropsTo(
      <Col xs={12} className={classes.trim()}>
        {this.props.children}
      </Col>
    );
  }
});

var PricingTableHeader = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='pricing-table-header text-center text-uppercase'>
        <h3>{this.props.children}</h3>
      </div>
    );
  }
});

var PricingTablePrice = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='pricing-table-price text-center text-uppercase'>
        <h2>{this.props.children}</h2>
      </div>
    );
  }
});

var PricingTableBody = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='pricing-table-body text-center'>
        {this.props.children}
      </div>
    );
  }
});

var PricingFeature = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='pricing-table-feature text-center'>
        {this.props.children}
      </div>
    );
  }
});


var PricingButtonContainer = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='pricing-btn-container text-center'>
        {this.props.children}
      </div>
    );
  }
});

module.exports.PricingTable = PricingTable;
module.exports.PricingFeature = PricingFeature;
module.exports.PricingTableBody = PricingTableBody;
module.exports.PricingTablePrice = PricingTablePrice;
module.exports.PricingTableHeader = PricingTableHeader;
module.exports.PricingTableContainer = PricingTableContainer;
module.exports.PricingButtonContainer = PricingButtonContainer;
