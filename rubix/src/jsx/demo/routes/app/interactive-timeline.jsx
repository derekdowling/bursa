/** @jsx React.DOM */

var Header = require('../../common/header.jsx');
var Sidebar = require('../../common/sidebar.jsx');
var Footer = require('../../common/footer.jsx');

var dataObject = JSON.parse(require('raw!./dataObject.json'));

var Body = React.createClass({
  componentDidMount: function() {
    createStoryJS({
      type:       'timeline',
      width:      '100%',
      height:     '600',
      id:         'my-timeline-story',
      source:     dataObject,
      embed_id:   'my-timeline'
    });
  },
  render: function() {
    return (
      <Container id='body'>
        <Grid>
          <Row>
            <Col sm={12}>
              <PanelContainer controlStyles='bg-orange65 fg-white'>
                <PanelHeader className='bg-orange65 fg-white'>
                  <Grid>
                    <Row>
                      <Col xs={12}>
                        <h3>Interactive Timeline</h3>
                      </Col>
                    </Row>
                  </Grid>
                </PanelHeader>
                <PanelBody>
                  <Grid>
                    <Row>
                      <Col xs={12}>
                        <div id='my-timeline' dir='ltr'></div>
                      </Col>
                    </Row>
                  </Grid>
                </PanelBody>
              </PanelContainer>
            </Col>
          </Row>
        </Grid>
        {this.props.children}
      </Container>
    );
  }
});

var classSet = React.addons.classSet;
var InteractiveTimeline = React.createClass({
  mixins: [SidebarMixin],
  render: function() {
    var classes = classSet({
      'container-open': this.state.open
    });
    return this.transferPropsTo(
      <Container id='container' className={classes}>
        <Sidebar />
        <Header />
        <Body>
          <Footer />
        </Body>
      </Container>
    );
  }
});

module.exports = InteractiveTimeline;
