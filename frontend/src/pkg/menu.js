import ax from "./axios"
import Admin from "../components/Admin.vue"
import Dashboard from "../components/Dashboard.vue"
import {isLogin} from "./auth"
import Login from "../components/Login.vue"
import NotFound from "../components/NotFound.vue"
import Menu from "../components/Menu.vue"
import {saveState} from "./session"

export default async function init(router) {
    if (!isLogin()) {
        router.addRoute({
            path: '/:pathMatch(.*)*',
            name: 'notFound',
            redirect: "/login"
        })
        router.addRoute({
            path: '/login',
            name: 'login',
            component: Login,
        })
        return
    }
    const menu = await ax({
        url: "system/menus",
    })
    if (!menu) {
        return
    }
    saveState({menus: menu})
    let routes = router.getRoutes()
    for (let i = 0; i < routes.length; i++) {
        router.removeRoute(routes[i].name)
    }
    router.addRoute({
        name: "login", path: "/login",
        component: Login
    })
    router.addRoute({
        path: '/:pathMatch(.*)*',
        name: 'notFound', component: NotFound
    })
    for (let i = 0; i < menu.length; i++) {
        let vc = toVueComponent(menu[i].component)
        if (vc) {
            menu[i].component = vc
            let au = menu[i].auth
            if (au) {
                menu[i].meta = {id: menu[i].id}
                menu[i].meta.auth = au
            }
            router.addRoute(menu[i])
        }
    }
}

function toVueComponent(com) {
    switch (com) {
        case "admin":
            return Admin
        case "dashboard":
            return Dashboard
        case "menu":
            return Menu
    }
    return undefined
}
