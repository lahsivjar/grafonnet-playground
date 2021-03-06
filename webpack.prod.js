const merge = require("webpack-merge");
const TerserPlugin = require('terser-webpack-plugin');
const common = require("./webpack.common.js");
const webpack = require("webpack");

module.exports = merge(common, {
  devtool: "source-map",
  plugins: [
    new webpack.DefinePlugin({
      "process.env.NODE_ENV": JSON.stringify("production")
    }),
  ],
  optimization: {
    minimizer: [new TerserPlugin()],
  },
});

