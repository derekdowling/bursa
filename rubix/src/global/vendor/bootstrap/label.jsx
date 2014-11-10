/** @jsx React.DOM */

var ColMixin = require('./col.jsx').ColMixin;

var Label = React.createClass({
  mixins: [ColMixin],
  propTypes: {
    inline: React.PropTypes.bool,
    control: React.PropTypes.bool
  },
  getLabelDOMNode: function() {
    return this.refs.label.getDOMNode();
  },
  render: function() {
    this.preRender();
    var classes = React.addons.classSet({
      'control-label': this.props.control,
      inline: this.props.inline
    });
    this.classes += ' ' + classes;
    return this.transferPropsTo(
      <label ref='label' className={this.classes} hidden={null}>
        {this.props.children}
      </label>
    );
  }
});

module.exports = Label;
