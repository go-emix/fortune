import {ElMessage} from 'element-plus';


export function Nerr(msg) {
    ElMessage({
        message: msg,
        type: 'error',
        duration: 2000
    })
}

export function Nsucc(msg) {
    ElMessage({
        message: msg,
        type: 'success',
        duration: 2000
    })
}