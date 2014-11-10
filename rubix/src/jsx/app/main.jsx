/** @jsx React.DOM */

'use strict';

/* Initialize Locales */
l20n.initializeLocales('app', {
  'locales': ['en-US'],
  'default': 'en-US'
});

/* Initializing touch events */
React.initializeTouchEvents(true);

/* APP PAGES */
var blank = require('./routes/app/blank.jsx');

/* ROUTES */
var routes = (
  <Routes>
    <Route name='root' path='/' view={blank} />
  </Routes>
);

Pace.once('hide', function() {
  $('#pace-loader').removeClass('pace-big').addClass('pace-small');
});

var InitializeRouter = function(View) {
  // cleanup
  if(window.Rubix) window.Rubix.Cleanup();
  Pace.restart();
  if(window.hasOwnProperty('ga') && typeof window.ga === 'function') {
    window.ga('send', 'pageview', {
     'page': window.location.pathname + window.location.search  + window.location.hash
    });
  }

  React.renderComponent(<View />, document.getElementById('app-container'), function() {
    // l20n initialized only after everything is rendered/updated
    l20n.ready();
    setTimeout(function() {
      $('body').removeClass('fade-out');
    }, 500);
  });
};

RRouter.routing = RRouter.HashRouting.start(routes, InitializeRouter);
