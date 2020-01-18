<!--
导入一个组件：import导入，然后在放到component中,然后就可以使用相应的标签了
选项对象：
  例如：
  {
    name: 'todoList',
    layout: '<template>一个todoList组件</template>'，
    data: {
      ... data被绑定到vue实例，并加入到响应式系统
    },
    props: {
      ...
    }
  }

  组件是一个Vue实例，new创建的是Vue跟实例。可以通过new Vue({el: "#component-demo"})的方式在创建根实例的时候传入一个组件。
  组件是可复用的Vue实例，与new Vue接受相同的选项对象。例如data、computed、watch、methods等。但是el是根实例才能有的。
  props，为组件提供属性。可以在v-bind使用 v-bind:tital="sd"

   编写模板时，每个组件必须只有一个根元素
   一个组件中的props可能很多，如果都使用bind一一指定会比较麻烦。但是props可以是一个对象，这利用了ES6的模板字符串特性。

  创建Vue实例时，需要指定一个模板，或者一个render函数。前者的情况是，Vue把模板编译成渲染函数。


使用Vue.component可以注册一个新的组件，传入一个组件名称和选项对象
组件的data必须是函数，用来保证是复用的
组件必须注册，使用Vue.component是全局注册，全局注册之后，任何通过new Vue创建的Vue实例都可以使用该组件。
全局注册必须在new Vue之前。
组件的命名，字母全小写且必须包含一个连字符。

局部注册：通过选项对象的components选项指定注册的组件。这样注册的组件的组件名就是自定义的组件名称。
这样注册之后，组件就可以在相应的局部范围使用了。

基础组件在不同的组件中可能会被频繁注册，为了避免这样。基础组件最好是全局注册。
--->

<template>
  <layout msg="Welcome to Your Vue.js App"/>
</template>

<script>
import layout from '../components/backLayout.vue'

export default {
  name: 'background_app',
  components: {
    layout
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
 body{
   margin: 0;
   box-sizing: border-box;
 }
</style>
