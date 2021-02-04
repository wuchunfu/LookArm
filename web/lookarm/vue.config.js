module.exports = {
  transpileDependencies: ['vuetify'],
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = 'LookArm - 快速搜索支持Mac M1的App'
      return args
    })
  },
  productionSourceMap: false
}
