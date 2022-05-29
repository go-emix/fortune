import {createRouter, createWebHashHistory, createWebHistory} from "vue-router"
import Dashboard from '../components/Dashboard.vue'
import Login from '../components/Login.vue'
import NotFound from '../components/NotFound.vue'
import Admin from '../components/Admin.vue'
import {isLogin, isPermit} from "./auth"
import {Nerr} from "./notify"
import i18n from "./i18n"
import {menu} from './menu'


function toVueComponent(com) {
    switch (com) {
        case "admin":
            return Admin
        case "login":
            return Login
        case "dashboard":
            return Dashboard
    }
}

const routes = [
    {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},
]

for (let i = 0; i < menu.length; i++) {
    menu[i].component = toVueComponent(menu[i].component)
    routes.push(menu[i])
}

const t = i18n.global.t

function routerHistory() {
    if (import.meta.env.DEV) {
        return createWebHistory()
    }
    return createWebHashHistory()
}

const router = createRouter({
    history: routerHistory(),
    routes,
});

router.beforeEach(function (to, from, next) {
    let au = to.meta.auth
    if (!au) {
        next()
        return
    }
    if (!isLogin()) {
        next({name: "login"})
        return
    }
    if (au === "login") {
        next()
        return
    }
    if (!isPermit(au)) {
        Nerr(t("not_permit"))
        return
    }
    next()
})

export default router