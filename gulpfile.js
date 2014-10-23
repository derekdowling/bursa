var gulp    = require('gulp')
, shell     = require('gulp-shell');

server_path = "src/bursa.io/server";
static_path = server_path + '/static/';

gulp.task('duo', shell.task([
    'duo ' + server_path + '/assets/build.js > ' + static_path + '/js/build.js',
    'duo --use duosass ' + server_path + '/assets/build.scss > ' + static_path + '/css/build.css'
]));

gulp.task('cjsx', shell.task([
    // 'cjsx-transform assets/coffee/*.coffee | coffee -cs > ' + server_path + '/assets/js/app.js'
]));

gulp.task('build', [
    'duo',
    'cjsx'
]);
