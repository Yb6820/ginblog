import axios from 'axios'
import Vue from 'vue'

axios.defaults.baseURL = 'http://121.37.246.78:80/api/v1'
Vue.prototype.$http = axios