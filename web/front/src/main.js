import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'

import moment from 'moment'

import './plugins/http'

//mdPreview
import VMdPreview from '@kangc/v-md-editor/lib/preview'
import '@kangc/v-md-editor/lib/style/preview.css'
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js'
import '@kangc/v-md-editor/lib/theme/style/github.css'

// 引入所有语言包
import hljs from 'highlight.js';

VMdPreview.use(githubTheme, {
  Hljs: hljs,
})

Vue.filter('dateformat', function (indate, outdate) {
  return moment(indate).format(outdate)
})

Vue.config.productionTip = false
Vue.use(VMdPreview)
new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app')
