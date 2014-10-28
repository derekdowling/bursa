var gulp    = require('gulp')
, shell     = require('gulp-shell');

// builds all compilable assets
gulp.task('build', [
    'js',
    'sass'
]);

gulp.task('js', shell.task([
    'cjsx -c --output static/js assets/coffee/marketing.cjsx',
    'cjsx -c --output static/js assets/coffee/app.cjsx',
    'duo --root static/js/ --output ./ app.js marketing.js'
]));     
   
gulp.task('sass', shell.task([
    'duo --use duosass --root assets/scss --output ../../static/css marketing.scss app.scss'
]));
