import {ElMessage} from 'element-plus';

const duration = 1500

export function Nerr(msg) {
    ElMessage({
        message: msg,
        type: 'error',
        duration
    })
}

export function Nsucc(msg) {
    ElMessage({
        message: msg,
        type: 'success',
        duration
    })
}