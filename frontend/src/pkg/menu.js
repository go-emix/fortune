import ax from "./axios"

let da = []

async function init() {
    da = await ax({
        url: "system/menu",
    })
}

init()

da = [
    {
        name: 'dashboard',
        component: "dashboard",
        path: '/',
        meta: {
            auth: "login"
        }
    },
    {
        name: 'admin',
        component: "admin",
        path: '/admin',
        meta: {
            auth: "admin",
            parent: "sys"
        }
    },
    {
        name: 'login',
        component: "login",
        path: '/login'
    },
]

export const menu = da
