import {getState} from "./session"

export function isLogin() {
    let state = getState();
    if (!state) {
        return false
    }
    return true
}

export function isPermit(to) {
    let state = getState();
    for (let i = 0; i < state.permit.length; i++) {
        let p = state.permit[i];
        if (p === "*" || p === to) {
            return true
        }
    }
    return false
}