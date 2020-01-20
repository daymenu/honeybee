module.exports = {
  devServer: {
    proxy: {
        '/ws/*' : {
            target: 'ws://localhost:8081',
            ws: true,
            changeOrigin: true
        }
    }
  }
}