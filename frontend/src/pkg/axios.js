import ax from 'axios'
import Cookies from "js-cookie";
import {Nerr, Nsucc} from './notify'
import i18n from './i18n'

const t = i18n.global.t

const env = import.meta.env

let instance = ax.create({
    baseURL: env.VITE_API_PATH,
    timeout: 2000,
})


instance.interceptors.request.use(function (config) {
    let to = Cookies.get("token")
    if (to) {
        config.headers["token"] = to
    }
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
        handleErr(code, data.reason)
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


function handleErr(errcode, reason) {
    switch (errcode) {
        case 1000:
            Nerr(reason)
            break
        case 1001:
            Nerr(reason)
            break
        default:
            Nerr(errcode)
    }
}
