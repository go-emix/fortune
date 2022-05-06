import ax from 'axios'

const env = import.meta.env

let instance = ax.create({
    baseURL: env.VITE_API_PATH,
    timeout: 2000,
})

instance.interceptors.request.use(function (config) {
    console.log(config)
}, function (error) {
    console.log(error)
})

instance.interceptors.response.use(function (resp){
    console.log(resp)
},function (error) {
    console.log(error)
})


export default instance

export function handleError(error) {
    if (!error.response) {
        alert("服务器连接中断")
    } else {
        console.log(error)
    }
}