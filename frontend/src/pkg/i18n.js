import {createI18n} from 'vue-i18n'

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

export default createI18n({
    locale: 'zh',
    messages: {
        en,
        zh
    }
})