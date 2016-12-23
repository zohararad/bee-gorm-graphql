const path = require('path'),
      projectRoot = path.resolve(__dirname),
      webpack = require('webpack'),
      CommonsChunkPlugin = require("webpack/lib/optimize/CommonsChunkPlugin"),
      ExtractTextPlugin = require('extract-text-webpack-plugin');

function assetsPath(_path) {
  return path.posix.join(projectRoot, 'static/', _path)
}

module.exports = [
  {
    name: 'js',
    entry: {
      app: './static/js/index.js',
      vendor: ['vue', 'vuex', 'vue-router', 'lokka-transport-http']
    },
    output: {
      path: path.join(projectRoot, 'public/assets'),
      publicPath: '/assets/',
      filename: '[name].js'
    },
    resolve: {
      extensions: ['', '.js', '.vue'],
      fallback: [path.join(__dirname, 'node_modules')],
      alias: {
        'vue$': 'vue/dist/vue.common.js',
        'src': path.resolve(__dirname, 'static/js'),
        'assets': path.resolve(__dirname, 'static/assets'),
        'components': path.resolve(__dirname, 'static/js/components')
      }
    },
    resolveLoader: {
      fallback: [path.join(__dirname, 'node_modules')]
    },
    module: {
      preLoaders: [
        {
          test: /\.vue$/,
          loader: 'eslint',
          include: projectRoot,
          exclude: /node_modules/
        },
        {
          test: /\.js$/,
          loader: 'eslint',
          include: projectRoot,
          exclude: /node_modules/
        }
      ],
      loaders: [
        {
          test: /\.vue$/,
          loader: 'vue'
        },
        {
          test: /\.js$/,
          loader: 'babel',
          include: projectRoot,
          exclude: /node_modules/
        },
        {
          test: /\.json$/,
          loader: 'json'
        },
        {
          test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
          loader: 'url',
          query: {
            limit: 10000,
            name: assetsPath('img/[name].[hash:7].[ext]')
          }
        },
        { test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "url-loader?limit=10000&minetype=application/font-woff" },
        { test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "file-loader" }
      ]
    },
    eslint: {
      formatter: require('eslint-friendly-formatter')
    },
    vue: {
      loaders: {
        css: ExtractTextPlugin.extract('vue-style-loader', "css-loader!sass-loader")
      },
      postcss: [
        require('autoprefixer')({
          browsers: ['last 2 versions']
        })
      ]
    },
    // eval-source-map is faster for development
    devtool: '#source-map',
    plugins: [
      new webpack.optimize.CommonsChunkPlugin('vendor.js'),
      new ExtractTextPlugin('components.css'),
      new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/)
    ]
  },
  {
    name: 'css',
    entry: {
      app: './static/css/index.scss'
    },
    output: {
      path: path.join(projectRoot, 'public/assets'),
      publicPath: '/assets/',
      filename: '[name].css'
    },
    module: {
      loaders: [
        { test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "url-loader?limit=10000&minetype=application/font-woff" },
        { test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "file-loader" },
        {
          test: /\.scss$/,
          loader: ExtractTextPlugin.extract('style-loader', "css-loader!sass-loader")
        }
      ]
    },
    devtool: '#source-map',
    plugins: [
      new ExtractTextPlugin('[name].css')
    ],
    resolve: {
      extensions: ['', '.scss'],
      root: [path.join(__dirname, 'static', 'css')]
    }
  }
];
