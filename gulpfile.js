// Import all plugins
const gulp    = require("gulp"),
      plumber = require("gulp-plumber"),
      notify  = require("gulp-notify"),
      babel   = require("gulp-babel"),
      watch   = require("gulp-watch")


// Store our config in one place
const paths = {
    js: {
        input:  "assets/js/**/*.js",
        output: "public/js"
    }
}

// Run js through babel so we can use es6 code.
gulp.task("js", () => {
    return watch(paths.js.input, {ignoreInitial: false})
        .pipe(plumber())
        .pipe(babel({
            presets: ["es2015"]
        }))
        .pipe(gulp.dest(paths.js.output))
        .pipe(notify("JS done processing"))
        .pipe(plumber.stop())
})


// Run it all.
gulp.task("default", ["js"])