// webpack.config.js
module.exports = {
    // ... other webpack configurations
    module: {
      rules: [
        // ... other rules
        {
          test: /\.module\.css$/,
          use: [
            'style-loader',
            {
              loader: 'css-loader',
              options: {
                modules: true,
              },
            },
          ],
        },
        // You might want to add a rule for non-module CSS files as well
        {
          test: /\.css$/,
          exclude: /\.module\.css$/,
          use: ['style-loader', 'css-loader'],
        },
      ],
    },
  };