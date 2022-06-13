import Cookies from "js-cookie";

export function saveState(obj) {
    let state = {}
    let sta = Cookies.get("state");
    if (sta) {
        state = JSON.parse(sta)
    }
    for (const k in obj) {
        let value = obj[k];
        switch (k) {
            case "token":
                state.token = value
                break
            case "i18n":
                state.i18n = value
                break
            case "menus":
                state.menus = value
                break
            case "components":
                state.components = value
                break
        }
    }
    Cookies.set("state", JSON.stringify(state))
}

export function getState() {
    let state = Cookies.get("state");
    if (!state) {
        return state
    }
    return JSON.parse(state);
}


export function clearState() {
    Cookies.remove("state")
}

export function exit() {
    clearState()
    location.reload()
}