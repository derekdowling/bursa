/** @jsx React.DOM */

var Dropdown = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='dropdown'>
        {this.props.children}
      </div>
    );
  }
});

module.exports = Dropdown;
