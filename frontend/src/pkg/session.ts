import Cookies from 'js-cookie';
import i18n from "./i18n";

export class Admin {
    id: number
    nickname: string
    token: string
    username: string

    constructor() {
        this.id = 0;
        this.nickname = ""
        this.token = ""
        this.username = ""
    }
}

export class Menu {
    auth: string
    component: string
    id: number
    name: string
    parent: number
    path: string

    constructor() {
        this.auth = ""
        this.component = ""
        this.id = 0
        this.name = ""
        this.parent = 0
        this.path = ""
    }
}

export class Feature {
    name: string
    id: number
    menu_id: number
    menu: Menu

    constructor() {
        this.name = ""
        this.menu = new Menu()
        this.id = 0
        this.menu_id = 0
    }
}

export class State {
    admin: Admin
    i18n: string
    menus: Menu[]
    features: Feature[]

    constructor() {
        this.admin = new Admin()
        this.i18n = "zh"
        this.menus = []
        this.features = []
    }
}


export function saveState(obj: any) {
    let state: State = new State()
    let sta = Cookies.get("state");
    if (sta) {
        state = JSON.parse(sta)
    }
    for (const k in obj) {
        let value = obj[k];
        switch (k) {
            case "admin":
                state.admin = value
                break
            case "i18n":
                state.i18n = value
                break
            case "menus":
                state.menus = value
                break
            case "features":
                state.features = value
                break
        }
    }
    Cookies.set("state", JSON.stringify(state))
}

export function getState(): State | null {
    try {
        let sta: string | undefined = Cookies.get("state");
        if (typeof sta === "string") {
            return JSON.parse(sta)
        }
    } catch (e) {
        exit()
    }
    return null
}

export function clearState(): void {
    Cookies.remove("state")
}

export function exit(): void {
    clearState()
    location.reload()
}