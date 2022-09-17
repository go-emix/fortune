import {createI18n} from 'vue-i18n'
import {getState} from "./session"
import ax from "axios"

const env = import.meta.env

let resp = await ax({
    url: env.VITE_API_PATH + "system/i18n"
})

let data = resp.data.data;

const en = data.en
const zh = data.zh

export const langs = [
    {
        i18n: en.i18n,
        lang: en.lang
    },
    {
        i18n: zh.i18n,
        lang: zh.lang
    },
]

function getLocale() {
    let state = getState();
    if (!state) {
        return 'zh'
    }
    if (state.i18n === "") {
        return 'zh'
    }
    return state.i18n
}

export default createI18n({
    locale: getLocale(),
    messages: {
        en,
        zh
    }
})