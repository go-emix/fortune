import {createRouter, createWebHashHistory, createWebHistory} from "vue-router"
import Dashboard from './components/Dashboard.vue'
import Login from './components/Login.vue'
import NotFound from './components/NotFound.vue'
import Admin from './components/Admin.vue'
import {isLogin, isPermit} from "./auth"
import {Nerr} from "./pkg/notify"
import i18n from "./pkg/i18n"

const t = i18n.global.t

const routes = [
    {
        name: 'dashboard',
        component: Dashboard,
        path: '/',
        meta: {
            auth: "login"
        }
    },
    {
        name: 'admin',
        component: Admin,
        path: '/admin',
        meta: {
            auth: "admin"
        }
    },
    {
        name: 'login',
        component: Login,
        path: '/login'
    },
    {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},
]

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

export default router;