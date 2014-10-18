var gulp    = require('gulp')
, shell     = require('gulp-shell');

static_path = "src/bursa.io/server/static/";

gulp.task('duo', shell.task([
    'duo index.js dist',
    'duo assets/scss/*.{scss} dist/css'
]));

gulp.task('cjsx', shell.task([
    'cjsx-transform assets/coffee/*.coffee | coffee -cs > ' + static_path + '/js/app.js';
]);

gulp.task('build', [
    'duo',
    'cjsx',
    'slim',
]);
