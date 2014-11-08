/** @jsx React.DOM */

var Caret = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <span className='caret'></span>
    );
  }
});

module.exports = Caret;
