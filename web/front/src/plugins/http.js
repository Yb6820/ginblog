import axios from 'axios'
import Vue from 'vue'

axios.defaults.baseURL = 'http://www.centyoubet.xyz/api/v1'
Vue.prototype.$http = axios