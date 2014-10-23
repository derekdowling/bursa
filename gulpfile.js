var gulp    = require('gulp')
, shell     = require('gulp-shell');

gulp.task('duo', shell.task([
    'duo assets/build.js > static/js/build.js',
    'duo --use duosass assets/build.scss > static/css/build.css'
]));

gulp.task('cjsx', shell.task([
    // 'cjsx-transform assets/coffee/*.coffee | coffee -cs > ' + server_path + '/assets/js/app.js'
]));

gulp.task('build', [
    'duo',
    'cjsx'
]);
