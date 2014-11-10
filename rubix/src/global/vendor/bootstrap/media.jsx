/** @jsx React.DOM */

var MediaBody = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='media-body'>
        {this.props.children}
      </div>
    );
  }
});

var Media = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <li className='media'>
        {this.props.children}
      </li>
    );
  }
});


var MediaDiv = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='media'>
        {this.props.children}
      </div>
    );
  }
});

var MediaList = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <ul className='media-list'>
        {this.props.children}
      </ul>
    );
  }
});

var MediaObject = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <img className='media-object' />
    );
  }
});

var MediaHeading = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <h4 className='media-heading fg-black50'>
        {this.props.children}
      </h4>
    );
  }
});

module.exports.Media = Media;
module.exports.MediaDiv = MediaDiv;
module.exports.MediaBody = MediaBody;
module.exports.MediaList = MediaList;
module.exports.MediaObject = MediaObject;
module.exports.MediaHeading = MediaHeading;
