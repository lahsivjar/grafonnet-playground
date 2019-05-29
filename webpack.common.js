const path = require("path");
const CleanWebpackPlugin = require("clean-webpack-plugin");

module.exports = {
    entry: {
        app: './ui/src/js/app.js',
    },
    output: {
        path: path.resolve(process.cwd(), 'dist'),
        filename: '[name].js'
    },
    plugins: [
        new CleanWebpackPlugin()
    ],
    module: {
        rules: [
            {
                test: path.join(__dirname, "./ui/src/js"),
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.css$/,
                use: [
                    "style-loader",
                    "css-loader"
                ]
            }
        ]
    }
};

