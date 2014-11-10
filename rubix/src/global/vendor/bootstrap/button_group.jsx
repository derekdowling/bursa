/** @jsx React.DOM */

var classSet = React.addons.classSet;
var ButtonGroup = React.createClass({
  propTypes: {
    lg: React.PropTypes.bool,
    sm: React.PropTypes.bool,
    xs: React.PropTypes.bool,

    vertical: React.PropTypes.bool,
    justified: React.PropTypes.bool,

    dropup: React.PropTypes.bool
  },
  render: function() {
    var isVertical = this.props.vertical ? true : false;
    var classes = classSet({
      'btn-group': !isVertical,

      'btn-group-lg': this.props.lg,
      'btn-group-sm': this.props.sm,
      'btn-group-xs': this.props.xs,

      'btn-group-vertical': isVertical,
      'btn-group-justified': this.props.justified,

      'dropup': this.props.dropup
    });
    return this.transferPropsTo(
      <div className={classes}>
        {this.props.children}
      </div>
    );
  }
});

module.exports = ButtonGroup;
