/** @jsx React.DOM */

var Tag = React.createClass({
  getDefaultProps: function() {
    return {
      href: '#',
      color: 'darkgreen45'
    };
  },
  render: function() {
    return (
      <RRouter.Link href={this.props.href} className={'left-tag border-hover-'+this.props.color+' bg-hover-'+this.props.color+' fg-hover-white bg-lightgray50 border-lightgray50 fg-text'}>{this.props.children}</RRouter.Link>
    );
  }
});

module.exports = Tag;
