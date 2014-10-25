var gulp    = require('gulp')
, shell     = require('gulp-shell');

// builds all compilable assets
gulp.task('build', [
    'site',
    'app'
]);

// builds the marketing site
gulp.task('site', [
    'site-js',
    'site-sass'
]);

gulp.task('site-js', shell.task([
    'cjsx -cb assets/coffee/marketing/main.cjsx > static/js/marketing-build.js',
    'duo static/js/marketing-build.js > static/js/marketing.js',
    'rm static/js/marketing-build.js'
]));     
   
gulp.task('site-sass', shell.task([
    'duo --use duosass assets/scss/marketing/main.scss > static/css/marketing.css'
]));

// builds the app
gulp.task('app', [
    'app-js',
    'app-sass'
]);

gulp.task('app-js', shell.task([
    'cjsx -cb assets/coffee/app/main.cjsx > static/js/app-build.js',
    'duo static/js/app-build.js > static/js/app.js',
    'rm static/js/app-build.js'
]));  

gulp.task('app-sass', shell.task([
    'duo --use duosass assets/scss/app/main.scss > static/js/app.js'
]));
