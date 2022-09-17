import {exit, getState} from "./session"

export function isLogin(): boolean | string {
    let state = getState();
    if (!state || !state.admin) {
        return false
    }
    return state.admin.token
}

export function isPermit(to: string): boolean {
    let menus = getState()!.menus;
    let ln = menus.length
    if (ln === 0) {
        exit()
    }
    for (let i = 0; i < ln; i++) {
        let m = menus[i]
        if (m.auth === to) {
            return true
        }
    }
    return false
}