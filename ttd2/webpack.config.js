/* eslint import/no-extraneous-dependencies: off */
const path = require('path')
const webpack = require('webpack')
const webpackBundleAnalyzer = require('webpack-bundle-analyzer')

const { env, port, publicPath } = {
  env: 'development',
  port: 8884,
  publicPath: 'http://localhost:8884',
  resourcesPath: 'cdn.bootcss.com',
}
const devPort = port - 1

const src = path.resolve(__dirname, './web')
const lib = path.resolve(__dirname, './node_modules')
const dst = path.resolve(__dirname, './public/static/assets')

const development = {
  devtool: 'eval-cheap-module-source-map',
  devServer: {
    port: devPort,
    https: false,
    compress: true,
    proxy: { '/': { target: `http://localhost:${port}/`, secure: false } },
  },
  context: `${src}`,
  resolve: {
    extensions: ['.js', '.jsx'],
    modules: [path.resolve(__dirname, './src'), lib],
  },
  output: {
    publicPath: `${publicPath}/`,
    filename: '[name]/[name].js',
    path: dst,
  },
  entry: {
    index: [`${src}/main.js`],
    common: [
      `${src}/lib.js`
    ],
  },
  module: {
    rules: [
      {
        include: src,
        test: /\.jsx?$/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              presets: [['env', { modules: false }], 'stage-0', 'react'],
            },
          },
        ],
      },
    ],
  },
  plugins: [
    new webpack.NamedModulesPlugin(),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify(env || 'development'),
    }),
    new webpack.optimize.CommonsChunkPlugin({
      name: 'common',
      filename: '[name]/[name].js',
      minChunks: Infinity,
    }),
  ],
}

if (env === 'production') {
  development.plugins.push(new webpackBundleAnalyzer.BundleAnalyzerPlugin())
}

const production = Object.assign({}, development, {
  devtool: 'source-map',
  devServer: undefined,
})

module.exports = Object.assign({}, { development }, { production })[env]
