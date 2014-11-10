var fs = require('fs');
var path = require('path');
var express = require('express');
var parse = require('ua-parser').parse;
var compression = require('compression');
var minify = require('html-minifier').minify;

var app = express();
app.use(compression());
app.use(express.static(path.join(process.cwd(), 'public')));

var package = require('./package.json');

var defaultAppName = process.env.APP ? process.env.APP : 'app';
var html = fs.readFileSync(path.join(process.cwd(), 'src', 'jsx', defaultAppName, 'index.html'), {
  encoding: 'utf8'
});

var createStyleTag = function(file, media) {
  media = media || 'screen';
  return "    <link media='"+media+"' rel='stylesheet' type='text/css' href='"+file+"'>\n";
};

var stylesheets = '';
if(process.env.NODE_ENV === 'development') {
  stylesheets += createStyleTag('/css/'+defaultAppName+'/raw/{dir}/main.css', 'screen,print');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/raw/{dir}/theme.css');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/raw/{dir}/colors.css');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/raw/{dir}/font-faces.css');
} else {
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/main-blessed1.css', 'screen,print');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/main.css', 'screen,print');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/theme.css');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/colors-blessed1.css');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/colors.css');
  stylesheets += createStyleTag('/css/'+defaultAppName+'/blessed/{dir}/font-faces.css');
}

html = html.replace(new RegExp('{app}', 'g'), defaultAppName);
html = html.replace(new RegExp('{stylesheets}', 'g'), stylesheets);
html = html.replace(new RegExp('{version}', 'g'), package.version);

var ltr = html.replace(new RegExp('{dir}', 'g'), 'ltr');
var rtl = html.replace(new RegExp('{dir}', 'g'), 'rtl');

app.get('/ltr/', function(req, res) {
  res.send(ltr);
});

app.get('/rtl/', function(req, res) {
  if(process.env.RTL === 'false')
    return res.redirect('/ltr/');
  res.send(rtl);
});

/** BEGIN X-EDITABLE ROUTES */

app.get('/xeditable/groups', function(req, res) {
  res.send([ 
    {value: 0, text: 'Guest'},
    {value: 1, text: 'Service'},
    {value: 2, text: 'Customer'},
    {value: 3, text: 'Operator'},
    {value: 4, text: 'Support'},
    {value: 5, text: 'Admin'}
  ]);
});

app.get('/xeditable/status', function(req, res) {
  res.send(500);
});

app.post('/xeditable/address', function(req, res) {
  res.send(200);
});

app.post('/dropzone/file-upload', function(req, res) {
  res.send(200);
});

/** END X-EDITABLE ROUTES */

var server = app.listen(process.env.PORT, function() {
  console.log('Server started on port %d', server.address().port);
});

process.on('uncaughtException', function() {
  process.exit(0);
});
