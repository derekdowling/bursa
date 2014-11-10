/** @jsx React.DOM */

var HelpBlock = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <p className='help-block'>
        {this.props.children}
      </p>
    );
  }
});

module.exports = HelpBlock;
