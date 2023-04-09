import axios from 'axios'
import Vue from 'vue'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
Vue.prototype.$http = axios
