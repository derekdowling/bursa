/** @jsx React.DOM */

var Row = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='row'>
        {this.props.children}
      </div>
    );
  }
});

module.exports = Row;
