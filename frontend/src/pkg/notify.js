import {ElMessage, ElMessageBox} from 'element-plus'
import i18n from './i18n'

const t = i18n.global.t

const duration = 1500
const duration2 = 2000

export function Nerr(msg) {
    ElMessage({
        message: msg,
        type: 'error',
        duration2
    })
}

export function Nsucc(msg) {
    ElMessage({
        message: msg,
        type: 'success',
        duration
    })
}

let style = {
    background: "#eef7f5",
    "border-radius": "8px",
    width: "250px",
    height: "130px",
}

export async function Nwarn(msg) {
    try {
        await ElMessageBox.confirm(
            msg,
            t('warning'),
            {
                confirmButtonText: t('ok'),
                cancelButtonText: t('cancel'),
                type: 'warning',
                customStyle: style,
            })
        return true
    } catch (e) {
        return false
    }
}
