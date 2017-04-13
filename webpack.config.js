var path = require("path");

module.exports = {
	entry: {
		bundle: ["babel-polyfill", "./frontend/js/main.js"]
	},
	output: {
		path: path.resolve(__dirname, "public/js"),
		publicPath: "/js/",
		filename: "[name].js"
	},
	module: {
		rules: [
			{
				test: /\.js$/,
				exclude: /node_modules/,
				loader: "babel-loader"
	 		}
		]
	},
	devtool: "source-map"
};
