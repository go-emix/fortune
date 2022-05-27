import {createRouter, createWebHistory} from "vue-router"
import Dashboard from './components/Dashboard.vue'
import Login from './components/Login.vue'
import NotFound from './components/NotFound.vue'
import {isLogin} from "./auth";

const routes = [
    {
        name: 'dashboard',
        component: Dashboard,
        path: '/',
        meta: {
            auth: isLogin
        }
    },
    {
        name: 'login',
        component: Login,
        path: '/login'
    },
    {path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound},
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach(function (to, from, next) {
    if (to.meta.auth) {
        let au = to.meta.auth()
        if (!au) {
            next({name: "login"})
        } else {
            next()
        }
    }
    next()
})

export default router;