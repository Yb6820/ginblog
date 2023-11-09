import axios from 'axios'
import Vue from 'vue'

axios.defaults.baseURL = 'http://106.53.113.95:80/api/v1'
Vue.prototype.$http = axios