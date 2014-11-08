/** @jsx React.DOM */

var ButtonToolbar = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='btn-toolbar' role='toolbar'>
        {this.props.children}
      </div>
    );
  }
});

module.exports = ButtonToolbar;
