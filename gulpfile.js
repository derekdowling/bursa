var del = require('del');
var path = require('path');
var gulp = require('gulp');
var shell = require('gulp-shell');

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
  scss: ['assets/css/app/**/*.scss', 'assets/css/marketing/**/.*scss'],
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

// Loads dependecies via duo
gulp.task('duo', function() {
    shell(['duo --use duosass --root assets --copy --output ../static css/vendor.scss js/build.js']);
});

// Doesn't work
gulp.task('watch', function () {
  gulp.watch('assets/**/*.scss', ['sass:marketing']);
});

/* ---------------------------------- */
/* --------- BEGIN MARKETING:SASS ---------- */
/* ---------------------------------- */
gulp.task('sass:marketing', function() {
  return gulp.src('./assets/css/marketing/main.scss')
          .pipe(sass({
            // sourceComments: 'normal' // uncomment when https://github.com/sass/node-sass/issues/337 is fixed
          }))
          .pipe(autoprefixer('last 2 versions', '> 1%', 'ie 9'))
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/marketing'));
});

gulp.task('minifycss:marketing', function() {
  return gulp.src(['static/css/marketing'])
          .pipe(minifycss())
          .pipe(gulp.dest('static/css/marketing'));
})

gulp.task('bless:marketing', function() {
  return gulp.src('static/css/marketing/*.css')
          .pipe(bless())
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/marketing/blessed'));
});

/* -------------------------------- */
/* --------- END MARKETING:SASS ---------- */
/* -------------------------------- */

/* ---------------------------------- */
/* --------- BEGIN APP:SASS ---------- */
/* ---------------------------------- */
gulp.task('sass:app', function() {
  return gulp.src('./assets/css/app/main.scss')
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
gulp.task('build:app', ['duo', 'sass:app']);
gulp.task('build:marketing', ['duo', 'sass:marketing']);

//This doesn't use the above tasks to avoid running the duo
//command multiple times per build
gulp.task('build', ['duo', 'sass:app', 'sass:marketing']);

// For Production Builds
gulp.task('bless', ['bless:app', 'bless:marketing']);
gulp.task('minifycss', ['minifycss:app', 'minifycss:marketing']);
gulp.task('build:dist', ['minifycss', 'bless']);

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
