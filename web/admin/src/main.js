import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/antui'
import './assets/css/style.css'
import './plugins/http'
//mdEditor
import VMdEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';

// highlightjs
import hljs from 'highlight.js';

VMdEditor.use(githubTheme, {
  Hljs: hljs,
});

Vue.config.productionTip = false
Vue.use(VMdEditor)
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
