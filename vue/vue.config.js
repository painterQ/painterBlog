/*vue.config.js
 *vue.config.js 是一个可选的配置文件，如果项目的 (和 package.json 同级的) 根目录中存在这个文件，
 *那么它会被 @vue/cli-service 自动加载。
 *你也可以使用 package.json 中的 vue 字段，但是注意这种写法需要你严格遵照 JSON 的格式来写。
 */
module.exports = {
    publicPath: '/',    //default "/"
    outputDir:'../static',   //default "/dist"
    assetsDir: 'public',
    //indexPath: 'background.html',    //相对于outputDir的路径
    filenameHashing: true,
    pages: {                    //多页面的配置
        background: {
            // page 的入口
            entry: 'src/page_background/main.js',
            // 模板来源
            template: 'src/page_background/index.html',
            // 在 dist/background.html 的输出
            filename: 'background.html',
            // 当使用 title 选项时，
            // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
            title: 'Background Page',
        },
        index :{
            entry: 'src/page_index/main.js',
            // 模板来源
            template: 'src/page_index/index.html',
            // 在 dist/background.html 的输出
            filename: 'index.html',
            // 当使用 title 选项时，
            // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
            title: 'Index Page',
        },
    },
};