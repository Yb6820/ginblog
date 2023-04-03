import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/antui'
import './assets/css/style.css'
import './plugins/http'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
