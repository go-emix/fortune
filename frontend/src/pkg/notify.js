import {ElMessage} from 'element-plus';

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