const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  lintOnSave: false,
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: (config) => {
    config.plugin('html').tap((args) => {
      args[0].title = '欢迎来到GinBlog'
      return args
    })
  },
  productionSourceMap: false,
})
