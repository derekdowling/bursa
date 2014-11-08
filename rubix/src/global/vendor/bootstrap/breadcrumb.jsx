/** @jsx React.DOM */

var Link = RRouter.Link;

var classSet = React.addons.classSet;
var BLink = React.createClass({
  propTypes: {
    active: React.PropTypes.bool
  },
  render: function() {
    var classes = classSet({
      'active': this.props.active
    });
    var children = (
      <Link href={this.props.href}>
        <span>{this.props.children}</span>
        <span> </span>
      </Link>
    );
    if(this.props.active)
      children = this.props.children;
    return this.transferPropsTo(
      <li className={classes} href={null}>
        {children}
      </li>
    );
  }
});

var Breadcrumb = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <ol className='breadcrumb'>
        {this.props.children}
      </ol>
    );
  }
});

module.exports.BLink = BLink;
module.exports.Breadcrumb = Breadcrumb;
