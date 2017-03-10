const path = require('path');
const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

var plugins = [
  new ExtractTextPlugin('[name]', {
      allChunks: true
  })
]

if (process.env.NODE_ENV == "production") {
  plugins.push(
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify('production')
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
        compress: {
            warnings: false
        }
    })
  )
}

//create entries
var pages = [
  'index'
]

var entries = {}
for (var i = 0; i < pages.length; i++){
  var fileName = pages[i]
  entries['js/' + fileName + '.js'] = path.resolve(__dirname, 'main', 'static', 'js', 'pages', fileName + '.jsx')
  entries['css/' + fileName + '.css'] = path.resolve(__dirname, 'main', 'static', 'sass', fileName + '.scss')
}

const compiler = {
  entry: entries,
  module: {
    loaders: [
      {
        exclude: /node_modules/,
        loader: 'babel',
        test: /\.(jsx|js)$/,
      },
      {
          test: /\.scss$/,
          loader: ExtractTextPlugin.extract('css!sass')
      },
    ],
  },
  output: {
    path: "./main/static/dist",
    filename: "[name]",
  },
  plugins: plugins
};

module.exports = compiler;
