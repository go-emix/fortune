import {createI18n} from 'vue-i18n'

import en from '../i18n/en.json'
import zh from '../i18n/zh.json'

export const langs = [
    {
        i18n: en.i18n.source,
        lang: en.lang.source
    },
    {
        i18n: zh.i18n.source,
        lang: zh.lang.source
    },
]

export default createI18n({
    locale: 'zh',
    messages: {
        en,
        zh
    }
})