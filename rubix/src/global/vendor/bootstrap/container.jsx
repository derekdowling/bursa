/** @jsx React.DOM */

var classSet = React.addons.classSet;
var Container = React.createClass({
  propTypes: {
    fluid: React.PropTypes.bool,
    fixed: React.PropTypes.bool
  },
  render: function() {
    var classes = classSet({
      'container': this.props.fixed,
      'container-fluid': this.props.fluid
    });
    return this.transferPropsTo(
      <div className={classes}>
        {this.props.children}
      </div>
    );
  }
});

module.exports = Container;
