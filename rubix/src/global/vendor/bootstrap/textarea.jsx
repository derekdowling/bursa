/** @jsx React.DOM */

var TextArea = React.createClass({
  getInputDOMNode: function() {
    return this.refs.textarea.getDOMNode();
  },
  getValue: function() {
    return this.getInputDOMNode().value;
  },
  render: function() {
    return this.transferPropsTo(
      <textarea ref='textarea' className='form-control' defaultValue={this.props.children} />
    );
  }
});

module.exports = TextArea;
