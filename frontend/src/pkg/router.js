import {createRouter, createWebHashHistory, createWebHistory} from "vue-router"
import {isLogin, isPermit} from "./auth"
import {Nerr} from "./notify"
import i18n from "./i18n"
import menu from './menu'
import feature from './feature'

const t = i18n.global.t

function routerHistory() {
    if (import.meta.env.DEV) {
        return createWebHistory()
    }
    return createWebHashHistory()
}

const router = createRouter({
    history: routerHistory(),
    routes: []
});

await menu(router)

await feature()

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