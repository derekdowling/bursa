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
