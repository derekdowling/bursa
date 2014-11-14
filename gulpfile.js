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
var browserSync = require('browser-sync');

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

var watchify = require('watchify');
var browserify = require('browserify');
var coffee_reactify = require('coffee-reactify');
var source = require('vinyl-source-stream');

var runSequence = require('run-sequence');

var package = require('./package.json');

var production = argv.production ? true : false;

/* file patterns to watch */
var paths = {
  // l20n: ['src/global/vendor/l20n/*.jsx'],
  scss: ['assets/css/app/**/*.scss'],
  // Ignore bourbon and bootstrap when watching because they're huge.
  watch_scss: ['assets/css/app/**/*.scss', '!assets/css/app/bootstrap-sass/**', '!assets/css/app/bourbon/**'],

  js: ['assets/js/**/*.js', 'assets/js/**/*.coffee', 'assets/js/**/*.cjsx'],
  entrypoints: ['./assets/js/app/app.coffee']
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
gulp.task('app:sass', function() {
  return gulp.src('./assets/css/app/main.scss')
          .pipe(sass({
            errLogToConsole: true
            // sourceComments: 'normal' // uncomment when https://github.com/sass/node-sass/issues/337 is fixed
          }))
          .pipe(autoprefixer('last 2 versions', '> 1%', 'ie 9'))
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/app'))
          .pipe(browserSync.reload({stream:true}));
});

gulp.task('app:minifycss', function() {
  return gulp.src(['static/css/app'])
          .pipe(minifycss())
          .pipe(gulp.dest('static/css/app'));
});

gulp.task('app:bless', function() {
  return gulp.src('static/css/app/*.css')
          .pipe(bless())
          .pipe(insert.prepend(banner()+'\n'))
          .pipe(insert.prepend('@charset "UTF-8";\n'))
          .pipe(gulp.dest('static/css/app/blessed'));
});

// VENDOR CSS
gulp.task('vendor:css',shell.task(['duo --use duosass --root assets --copy --output ../static css/vendor.scss']));

// APP JS
// If this breaks - it feels brittle there's always:
// watchify -t coffee-reactify -t reactify assets/js/app/app.coffee -o static/js/build.js -d
var bundlejs = function(watch) {
  return function() {
    var bundler, rebundle;

    bundler = browserify(paths.entrypoints, {
      basedir: path.join(__dirname),
      debug: true,
      cache: {}, // required for watchify
      packageCache: {}, // required for watchify
      fullPaths: watch, // required to be true only for watchify
    });

    if(watch) {
      bundler = watchify(bundler);
    }

    bundler.transform(coffee_reactify);

    rebundle = function() {
      return bundler.bundle()
        .on('error', gutil.log.bind(gutil,'Browserify error'))
        .pipe(source('build.js'))
        .pipe(gulp.dest('static/js'))
        .pipe(browserSync.reload({stream:true}));
    };

    bundler.on('update', rebundle);

    return rebundle();
  };
};

gulp.task('js:app', bundlejs(false));
gulp.task('js:app:watch', bundlejs(true));

// META TASKS

// CSS Related
gulp.task('sass', ['app:sass']);
gulp.task('minifycss', ['app:minifycss']);
gulp.task('bless', ['app:bless']);

gulp.task('app:css', ['app:sass']);

gulp.task('build:css', ['app:css', 'vendor:css']);
gulp.task('build:js', ['js:app']);
gulp.task('build', ['build:css', 'build:js']);

gulp.task('build:dist', ['minifycss', 'bless', 'uglify']);

// WATCHING
gulp.task('build:css:watch', ['app:css'], ready);
gulp.task('build:js:watch', ['js:app:watch'], ready);
gulp.task('react-bootstrap:watch', ['react-bootstrap'], ready);

gulp.task('watch:all', function() {
  gulp.watch(paths.watch_scss, ['build:css:watch']);
  gulp.watch(paths.js, ['build:js:watch']);
});

gulp.task('watch', ['watch:all', 'browser-sync']);

gulp.task('browser-sync', function() {
  browserSync({
    proxy: "localhost:8080"
  });
});
