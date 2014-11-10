/** @jsx React.DOM */

var Hero = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='homepage-hero'>
        <Container fixed>
          {this.props.children}
        </Container>
      </div>
    );    
  }
});

var HeroHeader = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='homepage-hero-header'>
        {this.props.children}
      </div>
    );
  }
});

var HeroHeader2 = React.createClass({
  render: function() {
    return this.transferPropsTo(
      <div className='homepage-hero-header2'>
        {this.props.children}
      </div>
    );
  }
});

var Homepage = React.createClass({
  handleNavigation: function(e) {
    $('body').addClass('fade-out');
    setTimeout(function() {
      RRouter.routing.navigate('/app/dashboard');
    }.bind(this), 250);
  },
  render: function() {
    return (
      <Container id='homepage-container'>
        <Button bsStyle='deepred' id='demo-btn' onClick={this.handleNavigation}>View Demo</Button>
        <div>
          <Hero className='text-center hidden-xs' style={{height: 475, backgroundImage: 'url(/imgs/homepage/background.png)', backgroundRepeat: 'no-repeat', backgroundSize: 'cover', overflow: 'hidden'}}>
            <img src='/imgs/homepage/simplepowerful.png' style={{marginTop: 5}} />
          </Hero>
          <Hero className='text-center visible-xs' style={{height: 270, backgroundImage: 'url(/imgs/homepage/background.png)', backgroundRepeat: 'no-repeat', backgroundSize: 'cover', overflow: 'hidden'}}>
            <img width='270' className='visible-xs' src='/imgs/homepage/simplepowerful.png' style={{margin: 'auto', marginTop: 30}} />
          </Hero>
        </div>
        <Hero>
          <HeroHeader>{"Don't wait for Web Components"}</HeroHeader>
          <HeroHeader2>{"Embrace React"}</HeroHeader2>
          <Grid>
            <Row>
              <Col sm={7} collapseLeft collapseRight>
                <p style={{marginTop: 60}}>
                  Facebook recently open-sourced React, a library for building User Interfaces. It uses a Virtual DOM implementation for ultra-high performance.
                </p>
                <p>
                  Rubix Admin app uses React for semantic markup coupled with CommonJS for composable Components. The result: <strong>clean and elegant code.</strong>
                </p>
              </Col>
              <Col sm={5} collapseLeft collapseRight>
                <div className='hidden-xs text-right'>
                  <img src='/imgs/homepage/reactcode.png' />
                </div>
                <div className='visible-xs text-center'>
                  <img width='250' src='/imgs/homepage/reactcode.png' />
                </div>
              </Col>
            </Row>
          </Grid>
        </Hero>
        <Hero>
          <HeroHeader2>{"Bootstrap on Steroids"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <div className='hidden-xs'>
              <img src='/imgs/homepage/bootstrapreact.png' />
            </div>
            <div className='visible-xs'>
              <img width='250' src='/imgs/homepage/bootstrapreact.png' />
            </div>
          </div>
          <p className='text-center'>
            Rubix implements custom React Components for Bootstrap enabling you to write shorter, semantic markup. Say Goodbye to unwieldy classnames and spaghetti code!
          </p>
        </Hero>
        <Hero>
          <HeroHeader>{"Internationalization and Localization"}</HeroHeader>
          <HeroHeader2>{"Mozilla L20n.js"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <div className='hidden-xs'>
              <img src='/imgs/homepage/mozflags.png' />
            </div>
            <div className='visible-xs'>
              <img width='250' src='/imgs/homepage/mozflags.png' />
            </div>
          </div>
          <p className='text-center'>
            {"Mozilla L20n is a developer friendly framework that places languages in the localizer's hand to create better translations. "}
          </p>
          <p className='text-center'>
            {"It removes the need for developers to thoroughly understand the specifics of a natural language and provides an opportunity for localizers to create better translations. Rubix ships with custom React component bindings for the framework."}
          </p>
        </Hero>
        <Hero>
          <HeroHeader2>{"Rubix Charts"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <div className='hidden-xs'>
              <img src='/imgs/homepage/rubixcharts.png' />
            </div>
            <div className='visible-xs'>
              <img width='250' src='/imgs/homepage/rubixcharts.png' />
            </div>
          </div>
          <p className='text-center'>
            Rubix Charts is an aesthetically beautiful, hand-crafted charting library created exclusively for Rubix Admin app. We used the awesome D3.JS library to write all the charting components (Line, Area, Stacked, Bar, Column, Pie and Donut) that power Rubix Charts.
          </p>
        </Hero>
        <Hero>
          <HeroHeader>{"Create complex layouts easily"}</HeroHeader>
          <HeroHeader2>{"Panels"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <div className='hidden-xs'>
              <img src='/imgs/homepage/panels.png' />
            </div>
            <div className='visible-xs'>
              <img width='250' src='/imgs/homepage/panels.png' />
            </div>
          </div>
          <p className='text-center'>
            {"Rubix Panels empowers developers to create complex layouts in addition to the awesome Grid provided by Twitter Bootstrap. Pretty much every example page showcased in the demo makes use of Panels for layout."}
          </p>
        </Hero>
        <Hero style={{position: 'relative', zIndex: 2}}>
          <HeroHeader>{"The Asset Pipeline"}</HeroHeader>
          <HeroHeader2>{"Gulp, Flip and Bless!"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <div className='hidden-xs'>
              <img src='/imgs/homepage/assetpipeline.png' />
            </div>
            <div className='visible-xs'>
              <img width='250' src='/imgs/homepage/assetpipeline.png' />
            </div>
          </div>
          <p className='text-center'>
            {"Gulp is a streaming build system. It's use of streams and code-over-configuration makes for a simpler and more intuitive build system."}
          </p>
          <p className='text-center'>
            {"Rubix's Asset Pipeline depends entirely on Gulp as its backbone. "}<strong>{"Everything is automated"}</strong>{": be it compiling JSX, SASS or even WebFonts! We make use of Twitter's "}<strong>{"css-flip"}</strong>{" for RTL support and the awesome "}<strong>{"blesscss"}</strong>{" library for fixing IE9 selectors and stylesheet bug."}
          </p>
        </Hero>
        <Hero className='subtle-bottom-shadow'>
          <HeroHeader>{"One Last Thing"}</HeroHeader>
          <HeroHeader2>{"Fanatical Support!"}</HeroHeader2>
          <div className='text-center' style={{marginTop: 25, marginBottom: 25}}>
            <img src='/imgs/homepage/support.png' />
          </div>
          <p className='text-center'>
            {"We have already provided extensive documentation on using/implementing Rubix. However, we take this a step further by ensuring version releases (which includes bug fixes, new features etc) for the next 6 months and general support for 1 year."}
          </p>
        </Hero>
        <div>
          <Hero className='text-center' style={{height: 215, backgroundImage: 'url(/imgs/homepage/background.png)', backgroundRepeat: 'no-repeat', backgroundSize: 'cover', overflow: 'hidden', backgroundPosition: '0% 100%'}}>
            <mq minWidth={550}>
              <h1 className='fg-white' style={{marginTop: 0, marginBottom: 25, fontWeight: 100}}>So what are you waiting for?</h1>
            </mq>
            <mq maxWidth={549}>
              <h3 className='fg-white' style={{marginTop: 0, marginBottom: 25, fontWeight: 100}}>So what are you waiting for?</h3>
            </mq>
            <Button lg outlined inverse retainBackground bsStyle='red' onClick={this.handleNavigation}>Click here to View Demo</Button>
          </Hero>
        </div>
      </Container>
    );
  }
});

module.exports = Homepage;
