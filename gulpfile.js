var gulp = require('gulp');
var shell = require('gulp-shell');

gulp.task('watch', function () {
  gulp.watch('assets/**/*.scss', ['css:main']);
});

gulp.task('css:vendor', function() {
    shell(['duo --use duosass --root assets --copy --output ../static css/vendor.scss']);
});

gulp.task('css:main',
  shell.task(
    'duo --use duosass --root assets --copy --output ../static css/app.scss css/marketing.scss js/build.js'
  )
);

// Slightly Stripped Rubix CSS

var del = require('del');
var path = require('path');
var gulp = require('gulp');

var flip = require('css-flip');
var map = require('map-stream');
var through = require('through');
var transform = require('vinyl-transform');
var child_process = require('child_process');

var argv = require('yargs').argv;

var sass = require('gulp-sass');
var gutil = require('gulp-util');
var bless = require('gulp-bless');
var insert = require('gulp-insert');
var concat = require('gulp-concat');
var rename = require('gulp-rename');
var replace = require('gulp-replace');
var uglify = require('gulp-uglifyjs');
var webpack = require('gulp-webpack');
var ttf2woff = require('gulp-ttf2woff');
var cssfont64 = require('gulp-cssfont64');
var minifycss = require('gulp-minify-css');
var autoprefixer = require('gulp-autoprefixer');

var runSequence = require('run-sequence');

var package = require('./package.json');

var production = argv.production ? true : false;

/* file patterns to watch */
var paths = {
  // index: ['src/jsx/'+defaultAppName+'/index.html', 'service.js'],
  // l20n: ['src/global/vendor/l20n/*.jsx'],
  jsx: ['src/jsx/*.jsx', 'src/global/requires/*.js', 'src/jsx/**/*.jsx', 'src/jsx/**/**/*.jsx', 'src/jsx/**/**/**/*.jsx', '!src/global/vendor/l20n/*.jsx', '!src/global/vendor/bootstrap/*.jsx'],
  scss: ['assets/css/transform/**/*.scss'],
  // bootstrap: ['src/global/vendor/bootstrap/*.jsx'],
  // ttf: ['public/fonts/dropbox/'+defaultAppName+'/*.ttf']
};

var banner = function() {
  return '/*! '+package.name+' - v'+package.version+' - '+gutil.date(new Date(), "yyyy-mm-dd")+
          ' [copyright: '+package.copyright+']'+' */';
};

function logData(data) {
  gutil.log(
    gutil.colors.blue(
      gutil.colors.bold(data)
    )
  );
}

logData('Environment : '+ (production ? 'Production':'Development'));

/* ---------------------------------- */
/* --------- BEGIN APP:SASS ---------- */
/* ---------------------------------- */
gulp.task('sass:app', function() {
  return gulp.src('./assets/css/transform/main.scss')
          .pipe(sass({
            // sourceComments: 'normal' // uncomment when https://github.com/sass/node-sass/issues/337 is fixed
          }))
          .pipe(autoprefixer('last 2 versions', '> 1%', 'ie 9'))
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/app'));
});

gulp.task('minifycss:app', function() {
  return gulp.src(['static/css/app'])
          .pipe(minifycss())
          .pipe(gulp.dest('static/css/app'));
});

gulp.task('bless:app', function() {
  return gulp.src('static/css/app/*.css')
          .pipe(bless())
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/app/blessed'));
});

/* -------------------------------- */
/* --------- END APP:SASS ---------- */
/* -------------------------------- */

/* ------------------------------ */
/* --------- GULP TASKS --------- */
/* ------------------------------ */
gulp.task('sass', ['sass:app']);
gulp.task('minifycss', ['minifycss:app']);
gulp.task('bless', ['bless:app']);

gulp.task('build:css', ['sass']);
gulp.task('build:app', ['app']);

gulp.task('build:dist', ['minifycss', 'bless', 'uglify']);

/*BEGIN: ALIASES FOR CERTAIN TASKS (for Watch)*/
gulp.task('build:app:watch', ['build:app'], ready);
gulp.task('build:css:watch', ['build:css'], ready);
gulp.task('react-bootstrap:watch', ['react-bootstrap'], ready);
/*END: ALIASES*/

gulp.task('watch', function() {
  // gulp.watch(paths.jsx, ['build:app:watch']);
  gulp.watch(paths.scss, ['rebuild:css']);
});

function ready() {
  gutil.log(
    gutil.colors.bgMagenta(
      gutil.colors.red(
        gutil.colors.bold('[          STATUS: READY          ]')
      )
    )
  );
}
