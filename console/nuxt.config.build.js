require('toml-require').install()
const env = require(`../nanoblog.tpl.toml`)

module.exports = {
  env,
  /*
  ** Headers of the page
  */
  head: {
    title: 'NanoBlog',
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {name: 'robots', content: 'none'},
    ],
    link: [
      {rel: 'icon', type: 'image/x-icon', href: '/api/console/favicon.ico'},
      { rel: 'manifest', href: `/api/console/manifest.json` }
    ],
  },
  /*
  ** Customize the progress bar color
  */
  loading: {color: '#4a4a4a'},
  css: [
    'vuetify/dist/vuetify.min.css',
    '~assets/scss/main.scss',
  ],
  plugins: [
    {src: '~/plugins/axios.js', ssr: false},
    {src: '~/plugins/init.js', ssr: false},
    {src: '~/plugins/nuxt-client-init.js', ssr: false},
  ],
  mode: 'spa',
  /*
  ** Build configuration
  */
  build: {
    publicPath: (env.StaticServer || env.Server) + '/console/dist/',
    extractCSS: true,
    ssr: false,
    /*
    ** Run ESLint on save
    */
    extend (config, ctx) {
      if (ctx.dev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/,
        })
      }
    },
  },
}
