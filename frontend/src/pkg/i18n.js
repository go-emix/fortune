import {createI18n} from 'vue-i18n'
import {getState} from "./session";

import en from '../i18n/en'
import zh from '../i18n/zh'

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
    if (state.i18n) {
        return state.i18n
    }
    return 'zh'
}

export default createI18n({
    locale: getLocale(),
    messages: {
        en,
        zh
    }
})