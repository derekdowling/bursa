/** @jsx React.DOM */

var classSet = React.addons.classSet;

var ListGroupItemText = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <p className='list-group-item-text'>{this.props.children}</p>
    );
  }
});

var ListGroupItemHeading = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <h4 className='list-group-item-heading'>{this.props.children}</h4>
    );
  }
});

var ListGroupItem = React.createClass({
  propTypes: {
    active: React.PropTypes.bool,
    disabled: React.PropTypes.bool,
    bsStyle: React.PropTypes.string,

    info: React.PropTypes.bool,
    danger: React.PropTypes.bool,
    warning: React.PropTypes.bool,
    success: React.PropTypes.bool
  },
  render: function() {
    var classesObj = {
      'list-group-item': true,
      'active': this.props.active,
      'disabled': this.props.disabled,
      'list-group-item-info': this.props.info,
      'list-group-item-danger': this.props.danger,
      'list-group-item-warning': this.props.warning,
      'list-group-item-success': this.props.success
    };

    if(this.props.bsStyle) {
      var bsStyles=this.props.bsStyle.split(',');
      for(var i=0; i < bsStyles.length; i++) {
        classesObj['list-group-'+bsStyles[i].trim()] = true;
      }      
    }

    var classes = classSet(classesObj);

    return this.transferPropsTo(
      <RRouter.Link href='#' className={classes}>
        {this.props.children}
      </RRouter.Link>
    );
  }
});

var ListGroup = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='list-group'>
        {this.props.children}
      </div>
    );
  }
});

module.exports.ListGroup = ListGroup;
module.exports.ListGroupItem = ListGroupItem;
module.exports.ListGroupItemText = ListGroupItemText;
module.exports.ListGroupItemHeading = ListGroupItemHeading;
