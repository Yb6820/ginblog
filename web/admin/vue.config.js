const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
    lintOnSave: false,
    transpileDependencies: true,
    publicPath: '/admin/',
    outputDir: 'dist',
    assetsDir: 'static'
})