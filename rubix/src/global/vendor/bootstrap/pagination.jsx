/** @jsx React.DOM */

var Link = RRouter.Link;

var classSet = React.addons.classSet;
var Page = React.createClass({
  propTypes: {
    active: React.PropTypes.bool,
    disabled: React.PropTypes.bool,
    begin: React.PropTypes.bool,
    end: React.PropTypes.bool,
    next: React.PropTypes.bool,
    previous: React.PropTypes.bool
  },
  getInitialState: function() {
    return {
      active: this.props.active || false,
      disabled: this.props.disabled || false
    };
  },
  render: function() {
    var classes = classSet({
      'next': this.props.next,
      'active': this.state.active,
      'disabled': this.state.disabled,
      'previous': this.props.previous
    });
    var children = null;
    if(this.props.begin) {
      children = '«';
    } else if(this.props.end) {
      children = '»';
    } else if(this.props.next) {
      children = <span>{this.props.children} →</span>
    } else if(this.props.previous) {
      children = <span>← {this.props.children}</span>
    } else if(this.props.active) {
      children = <span>{this.props.children}<span className='sr-only'>(current)</span></span>
    } else {
      children = this.props.children
    }
    return this.transferPropsTo(
      <li href={null} className={classes}>
        <Link href={this.props.href || '#'}>
          {children}
        </Link>
      </li>
    );
  }
});

var Pagination = React.createClass({
  propTypes: {
    lg: React.PropTypes.bool,
    sm: React.PropTypes.bool
  },
  render: function() {
    var classes = classSet({
      'pagination': true,
      'pagination-lg': this.props.lg,
      'pagination-sm': this.props.sm
    });
    return this.transferPropsTo(
      <ul className={classes}>
        {this.props.children}
      </ul>
    );
  }
});

var Pager = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <ul className='pager'>
        {this.props.children}
      </ul>
    );
  }
});

module.exports.Page = Page;
module.exports.Pager = Pager;
module.exports.Pagination = Pagination;
