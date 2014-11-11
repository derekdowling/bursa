var gulp = require('gulp');
var shell = require('gulp-shell');

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

var cjsx = require('gulp-cjsx');
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
  // l20n: ['src/global/vendor/l20n/*.jsx'],
  jsx: ['src/jsx/*.jsx', 'src/global/requires/*.js', 'src/jsx/**/*.jsx', 'src/jsx/**/**/*.jsx', 'src/jsx/**/**/**/*.jsx', '!src/global/vendor/l20n/*.jsx', '!src/global/vendor/bootstrap/*.jsx'],
  cjsx: ['assets/cjsx/**/*.cjsx'],
  scss: ['assets/css/transform/**/*.scss']
};

// UTILITY
function banner()
{
  return '/*! '+package.name+' - v'+package.version+' - '+gutil.date(new Date(), "yyyy-mm-dd")+
          ' [copyright: '+package.copyright+']'+' */';
}

function logData(data) {
  gutil.log(
    gutil.colors.blue(
      gutil.colors.bold(data)
    )
  );
}

function ready() {
  gutil.log(
    gutil.colors.bgMagenta(
      gutil.colors.red(
        gutil.colors.bold('[          STATUS: READY          ]')
      )
    )
  );
}

logData('Environment : '+ (production ? 'Production':'Development'));

// SASS TASKS
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

// VENDOR CSS
gulp.task('css:vendor', function() {
    shell(['duo --use duosass --root assets --copy --output ../static css/vendor.scss']);
});

// CJSX
gulp.task('cjsx', function() {
  gulp.src(paths.cjsx)
    .pipe(cjsx({bare: true}).on('error', gutil.log))
    .pipe(concat('views.js'), gutil.log)
    .pipe(gulp.dest('static/views'), gutil.log);
});

gulp.task('uglify:cjsx', function() {
  gulp.src('static/views/*.js')
    .pipe(cjsx({bare: true}).on('error', gutil.log))
    .pipe(gulp.dest('static/views'), gutil.log);
});

// APP JS
gulp.task('js:app',
  shell.task(
    'duo --root assets --copy --output ../static js/app/app.coffee'
  )
);

// META TASKS

// CSS Related
gulp.task('sass', ['sass:app']);
gulp.task('minifycss', ['minifycss:app']);
gulp.task('bless', ['bless:app']);

// JS Related
gulp.task('uglify', ['uglify:cjsx']);

gulp.task('build:css', ['sass']);
gulp.task('build:js', ['cjsx', 'js:app']);
gulp.task('build', ['build:css', 'build:js']);

gulp.task('build:dist', ['minifycss', 'bless', 'uglify']);

// WATCHING
gulp.task('build:css:watch', ['build:css'], ready);
gulp.task('build:js:watch', ['build:js'], ready);
gulp.task('react-bootstrap:watch', ['react-bootstrap'], ready);

gulp.task('watch', function() {
  gulp.watch(paths.scss, ['build:css:watch']);
  gulp.watch(paths.cjsx, ['build:js:watch']);
});
