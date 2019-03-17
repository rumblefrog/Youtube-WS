const path = require('path');

module.exports = {
    pages: {
        index: {
            entry: 'main.ts',
        }
    },
    configureWebpack: {
        resolve: {
            alias: {
                "@": path.resolve('.'),
            }
        }
    },
    integrity: true,
    publicPath: undefined,
    outputDir: undefined,
    assetsDir: undefined,
    runtimeCompiler: undefined,
    productionSourceMap: undefined,
    parallel: undefined,
    css: undefined
}