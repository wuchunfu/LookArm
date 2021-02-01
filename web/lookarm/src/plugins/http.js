import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'https://lookarm.cn/api/v1'
Vue.prototype.$http = axios
