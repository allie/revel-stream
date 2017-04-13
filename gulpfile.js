var gulp = require("gulp");
var gutil = require("gulp-util");
var sass = require("gulp-sass");
var postcss = require("gulp-postcss");
var postcssImport = require("postcss-import");

gulp.task("assets", function() {
	gulp.src("frontend/assets/**/*.png")
		.pipe(gulp.dest("public/assets"));
});

gulp.task("styles", function() {
	gulp.src("frontend/sass/main.scss")
		.pipe(sass().on("error", sass.logError))
		.pipe(postcss([postcssImport()]))
		.pipe(gulp.dest("public/css"));
});

gulp.task("default", ["styles", "assets"]);

gulp.task("watch", function() {
	gulp.watch("frontend/sass/**/*.scss", ["sass"])
	gulp.watch("frontend/assets/**/*.png", ["assets"])
});
