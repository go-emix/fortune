import {exit, getState} from "./session"

export function isLogin() {
    let state = getState();
    if (!state) {
        return false
    }
    try {
        return state.admin.token
    } catch (e) {
        exit()
    }
}

export function isPermit(to) {
    let menus = getState().menus;
    if (!menus) {
        exit()
    }
    let ln = menus.length
    for (let i = 0; i < ln; i++) {
        let m = menus[i]
        if (m.auth === to) {
            return true
        }
    }
    return false
}