import axios from 'axios'

const instance = axios.create({
  // baseURL: 'http://localhost:8086',
  headers: {
    'Content-Type': 'application/json'
  }
})


// 请求拦截器
instance.interceptors.request.use(config => {
  // 在请求发送之前处理请求的loading状态等
  return config
}, error => {
  // 处理请求错误
  return Promise.reject(error)
})

// 响应拦截器
instance.interceptors.response.use(response => {
  // 处理响应的loading状态等
  return response
}, error => {
  // 处理响应错误
  return Promise.reject(error)
})

export default instance