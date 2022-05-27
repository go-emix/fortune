import {getState} from "./pkg/session"

export function isLogin() {
    let state = getState();
    if (!state) {
        return false
    }
    return true
}

export function isPass() {
    if (!isLogin()) {
        return false
    }
    return true
}