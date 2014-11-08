/** @jsx React.DOM */

var Footer = React.createClass({
  getInitialState: function() {
    return {
      version: 0
    };
  },
  componentDidMount: function() {
    this.setState({
      version: document.getElementsByTagName('body')[0].getAttribute('data-version')
    });
  },
  render: function() {
    return (
      <div>
        <Grid gutterBottom></Grid>
        <Grid id='footer' className='text-center'>
          <Row>
            <Col xs={12}>
              <div>© 2014 SketchPixy Creative - v{this.state.version}</div>
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
});

module.exports = Footer;
