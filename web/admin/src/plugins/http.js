import axios from 'axios'
import Vue from 'vue'

const Url = 'http://www.centyoubet.xyz/api/v1/'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})

Vue.prototype.$http = axios
export { Url }

