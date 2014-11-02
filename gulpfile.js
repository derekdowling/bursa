var gulp = require('gulp');
var shell = require('gulp-shell');

gulp.task('watch', function () {
  gulp.watch('assets/**/*.scss', shell.task([
    'duo --use duosass --root assets --output ../static css/app.scss css/marketing.scss js/build.js'
  ]));
});

gulp.task('css:vendor', function() {
    shell(['duo --use duosass --root assets --output ../static css/vendor.scss']);
});
