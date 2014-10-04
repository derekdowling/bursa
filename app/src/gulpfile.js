var gulp = require('gulp')
,   sass = require('gulp-sass')
,   minifyCSS = require('gulp-minify-css')
,   concat = require('gulp-concat')
,   plumber = require('gulp-plumber')
,   uglify = require('gulp-uglify')
,   rename = require('gulp-rename')
,   del = require('del')
,   coffee = require('gulp-coffee')
,   gutil = require('gulp-util')
,   watch = require('gulp-watch')
,   sourcemaps = require('gulp-sourcemaps');

var onError = function(err) {
    console.log(err);
}

// var bowerVendorDir = './src/Donata/AppBundle/Resources/public/vendor';
var bowerVendorDir = './bower_components';
var appPublicDir = './src/Donata/AppBundle/Resources/public';

var appPublicJS = appPublicDir + '/dist/js';
var appPublicCSS = appPublicDir + '/dist/css';

gulp.task('clean', function(cb) {
    del([
        appPublicJS,
        appPublicCSS
    ], cb);
});

gulp.task('import_js_libs', function() {
    gulp
        .src(bowerVendorDir + '/jquery/dist/jquery.min.js')
        .pipe(gulp.dest(appPublicJS));
    gulp
        .src(bowerVendorDir + '/jquery/dist/jquery.min.map')
        .pipe(gulp.dest(appPublicJS));
    gulp
        .src(bowerVendorDir + '/bootstrap/dist/js/bootstrap.min.js')
        .pipe(gulp.dest(appPublicJS));
    gulp
        .src(bowerVendorDir + '/d3/d3/min.js')
        .pipe(gulp.dest(appPublicJS));
});

gulp.task('import_css_libs', function() {
    gulp
        .src(bowerVendorDir + '/bootstrap/dist/css/bootstrap.min.css')
        .pipe(gulp.dest(appPublicCSS));

    gulp
        .src(bowerVendorDir + '/bootstrap/dist/css/bootstrap-theme.min.css')
        .pipe(gulp.dest(appPublicCSS));
});

gulp.task('import_assets', function() {
    gulp
        .src(bowerVendorDir + '/bootstrap/dist/fonts/bootstrap/*.{ttf,woff,eof,svg,eot}')
        .pipe(gulp.dest(appPublicDir + '/fonts/bootstrap/'));
});

gulp.task('compile_sass', function() {
    gulp
        .src(appPublicDir + '/scss/*.scss')
        .pipe(sass())
        .pipe(minifyCSS().on('error', gutil.log))
        .pipe(rename({suffix: ".min"}))
        .pipe(gulp.dest(appPublicCSS));
});

gulp.task('compile_coffee', function() {
    gulp
        .src(appPublicDir + '/coffee/*.coffee')
        .pipe(sourcemaps.init())
        .pipe(coffee().on('error', gutil.log))
        .pipe(sourcemaps.write())
        .pipe(minifyCSS())
        .pipe(rename({suffix: ".min"}))
        .pipe(gulp.dest(appPublicJS));
});

gulp.task('build', [
    'import_js_libs',
    'import_css_libs',
    'import_assets',
    'compile_sass',
    'compile_coffee'
]);

// NEEDS TWEEKING
// gulp.task('watch', function() {
    // watch(appPublicDir + 'coffee/*.coffee', function(files, cb) {
        // gulp.start('compile_coffee');
    // });
    // watch(appPublicDir + 'scss/*.scss', function(files, cb) {
        // gulp.start('compile_sass');
    // });
// });
