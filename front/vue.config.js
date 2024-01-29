// const { defineConfig } = require('@vue/cli-service')
// const path = require('path')
module.exports = {
  transpileDependencies: true,
  productionSourceMap: false,
  publicPath :process.env.VUE_APP_BASE_URL,
  devServer: {
    // host:"172.17.8.75",
    port:"8081",
    proxy: {
        '/apis': {
            target: 'http://127.0.0.1:8086/apis', //接口域名
            changeOrigin: true,             //是否跨域
            // ws: true,                       //是否代理 websockets
            secure: false,                   //是否https接口
            pathRewrite: {                  //路径重置
                '^/apis': ''
            },
            headers: {
                'Access-Control-Allow-Origin': '*' // 允许所有来源跨域访问
              }
        }
    }
}
}
