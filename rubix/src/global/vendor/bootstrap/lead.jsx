/** @jsx React.DOM */

var Lead = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <p className='lead'>
        {this.props.children}
      </p>
    );
  }
});

module.exports = Lead;
