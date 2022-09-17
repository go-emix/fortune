import ax from 'axios'
import {Nerr, Nsucc} from './notify'
import i18n from './i18n'
import {exit} from "./session"
import {isLogin} from "./auth"

const t = i18n.global.t

const env = import.meta.env

let instance = ax.create({
    baseURL: <string>env.VITE_API_PATH,
    timeout: 2500,
})

instance.interceptors.request.use(function (config) {
    let token = isLogin()
    if (token) {
        config.headers!["token"] = token
    }
    config.headers!["lang"] = String(i18n.global.locale)
    return config
}, function (error) {
    Nerr(error)
})

instance.interceptors.response.use(function (resp) {
    let data = resp.data
    if (!data) {
        Nerr(t("req_failed"))
        return
    }
    let code = data.errcode
    if (code !== 0) {
        if (code === 1004) {
            Nerr(code + " : " + data.errmsg)
            setTimeout(function () {
                exit()
            }, 1500)
            return
        } else if (code === 1005) {
            let url = resp.config.url
            let method = resp.config.method
            let msg = method + " " + url
            Nerr(code + " : " + msg + " " + data.errmsg)
            return
        }
        Nerr(code + " : " + data.errmsg)
        return
    }
    let rdata = data.data
    if (rdata) {
        return Promise.resolve(rdata)
    } else {
        let msg = data.errmsg
        Nsucc(t(msg))
        return Promise.resolve(msg)
    }
}, function () {
    Nerr(t("server_not_connected"))
})

export default instance