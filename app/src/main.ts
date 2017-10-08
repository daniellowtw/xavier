// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import { App } from './App'
import * as Vue from 'vue'
import * as VueResource from 'vue-resource'
import store from './store'

require('font-awesome/css/font-awesome.css')
Vue.config.productionTip = false
Vue.use(VueResource)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  store,
  components: { App }
})
