# painterBlog

## Project setup
```
初始化 ：vue create --default .
    通过命令行可以指定预设presets

启动： npm run serve --scripts-prepend-node-path=auto

cli-plugin-bable 
    可以通过配置文件babel.config.js修改
vue-cli-service 
    使用命令 vue-cli-service lint 
    可以通过.eslintrc或者package.json的eslintConfig域配置
    可以通过vue add eslint安装
package.json
    npm的配置文件，记录了了依赖

Babel
    一个js编译器，实现语法转换等
webpackage
    一个打包工具，
    vue.config.js 可以配置vue项目
        https://github.com/vuejs/vue-cli/tree/dev/docs/config
        指定打包后的输出目录
        保存时是否使用eslint-loader 来格式化
        webpack配置
        vue-loader的配置
        css-loader的配置
        webpack dev server相关的配置，例如端口等
        ...
    主要是：@vue/cli-service 依赖 webpack
    
配置一个vue3.0项目
    https://blog.csdn.net/moonbc/article/details/89549154

```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

### WebPack
入门文档：
https://www.webpackjs.com/guides/
```
npm init -y
npm install webpack webpack-cli --save-dev
注意修改package.json 增加+   "private": true,，以免意外发布代码
配置文件是webpack.config.js
    配置文件中可以指明enrty，和输出文件的名称（一般进入点称为main.js,输出称为bundle.js）
    指定dist的路径
    一般默认在package.js中指定"build": "webpack",这样npm run build就可以了
    
传统不使用打包工具的话，存在如下问题：
    不能直观看出依赖关系
    脚本引入的顺序有要求，容易出错
    可能导致浏览器下载无用代码

分出/src目录和/dist目录
安装包，只需要使用
    npm install --save name
    npm install --save-dev name

模块化后，使用import...from导入，这样明确依赖关系
不需要再index.html中增加一堆的<script>标签引入依赖，只需要引入根上的一个js即可
这个js就是打包的结果，一般命名为main.js

另外除了打包，webpack还会对代码进行压缩，使用loader系统进行转义处理，比如babel

```
可以加载css loader
```
npm install --save-dev style-loader css-loader
这样在webpack.config.js中会多出一个rules。
其中指定了对哪些文件（正则匹配文件名）使用哪些loader，比如这里使用style-loader和css-loader

在main.js中指定import相应的css文件，然后就会加到最后的bundle.js中
```
加载图片
```
npm install --save-dev file-loader
使用import Icon from './icon.png';导入图片，这个Icon实际是一个转化后的url
file-loader 和 url-loader 可以接收并加载任何文件，包括字体（使用@font-face导入）

还有 csv-loader 和 xml-loader ，直接得到解析好的json
```

### @vue/cli
当前目录自动推导入口文件——入口可以是 main.js、index.js、App.vue 或 app.vue 中的一个

### promise语法
1.解决回调地狱（Callback Hell）问题，使得多层嵌套回调变的扁平，称为类似链式调用的形式
不断的调用promise对象的then方法来完成回调的嵌套（前一个方法要返回需要回调的函数）

2. 方便在上述场景下捕获异常


### 垂直铺满屏幕（viewpoint）的两种方案
1.html元素设置100%

2.高度使用vh单位

### Vue Router 笔记
通过组合组件来组成应用程序，当你要把 Vue Router 添加进来，我们需要做的是，
0.导入Vue和Vue Router，使用Vue.use(VueRouter)
1.***将组件 (components) 映射到路由 (routes)***
2.告诉 Vue Router 在哪里渲染它们。

创建一个Router实例，然后作为选项对象的属性传入根Vue
创建Router实例时，需要传入一个路由配置数组，有path和component

### Vue模板语法:插值、指令、修饰符 :[aaa].modifier="bbb"
支持任何js表达式，不论是指令中还是Mustache语法中。但是这是一个沙盒环境，不能访问外部的全局变量，除了白名单中的Date和Math。
v-html,不能复合使用模板
v-once
v-bind  ===> :
v-on ===> @
v-if
HTML语法对指令的影响：HTML要求attr的名称不区分大小写，而且不能有字符：比如空格和引号
    属性名称应该使用小写，属性值永远加引号（双引号或者单引号）（所以自定义的bind应该返回字符串）
动态指令，可以使用```v-bind:[attributeName]="url"```的语法动态改变指令，attributeName也应该是一个js表达式，但值必须是字符串或null
指令修饰符
    指令后可以增加修饰符，例如```v-on:submit.prevent="onSubmit"```
    prevent
### Vue基础
数据监听（响应式的数据）
事件监听
自定义事件和自己触发事件
    $emit('事件名',抛出值)
    事件可以携带抛出值
双向绑定
    v-module = v-bind + v-on
    bind的是value属性，on的是input事件，事件处理是把value设置为事件抛出值
   
动态组件 
    <component> + is
    <keep-alive> 缓存失活的组件
异步组件
插槽和内容分发
    如果自定义组件没有包含一个 <slot> 元素，则该组件起始标签和结束标签之间的任何内容都会被抛弃。
    因为Vue不知道该把内容放到解析后模板的哪里

### 渲染函数
### 实例属性
$data, vm代理
$props,vm代理
$el
$options
$parent 父实例,可能没有
$root 根实例，可能是自己
$children
$slots
$ref
  


index的Store初始化的时候会生成一个docList实例，实例初始化的时候会请求mate信息。这应该对应目录的第一页。
{start: "/doc0", length: 10}

DocList初始化的时候也能指定一些init，但是现在是空的。
因此初始状态，DocList是空的, docsUpdate = flase, total=0

目前一开始获取全部的mate