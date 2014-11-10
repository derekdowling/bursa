/** @jsx React.DOM */

var Static = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <p className='form-control-static'>
        {this.props.children}
      </p>
    );
  }
});

module.exports = Static;
