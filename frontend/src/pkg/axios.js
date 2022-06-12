import ax from 'axios'
import {Nerr, Nsucc} from './notify'
import i18n from './i18n'
import {clearState, getState} from "./session";

const t = i18n.global.t

const env = import.meta.env

let instance = ax.create({
    baseURL: env.VITE_API_PATH,
    timeout: 2500,
})

instance.interceptors.request.use(function (config) {
    let state = getState();
    if (state) {
        config.headers["token"] = state.token
    }
    config.headers["lang"] = String(i18n.global.locale)
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
        Nerr(code + " : " + data.errmsg)
        if (code === 1003) {
            setTimeout(function () {
                clearState()
                location.reload()
            }, 1500)
        }
        return
    }
    let rdata = data.data
    if (rdata) {
        return Promise.resolve(rdata)
    } else {
        let msg = data.errmsg
        Nsucc(msg)
        return Promise.resolve(msg)
    }
}, function () {
    Nerr(t("server_not_connected"))
})

export default instance