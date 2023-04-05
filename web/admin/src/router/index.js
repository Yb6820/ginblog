import Vue from 'vue'
import VueRouter from 'vue-router'
import Admin from '../views/Admin.vue'
import Login from '../views/Login.vue'

// 页面路由组件
import Index from '../components/admin/Index.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'

Vue.use(VueRouter)

const routes = [{
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/',
        name: 'admin',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: Admin,
        children: [
            { path: 'index', component: Index },
            { path: 'addart', component: AddArt },
            { path: 'addart/:id', component: AddArt, props: true },
            { path: 'artlist', component: ArtList },
            { path: 'catelist', component: CateList },
            { path: 'userlist', component: UserList }
        ]
    }
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, from, next) => {
    const token = window.sessionStorage.getItem('token')
    if (to.path === '/login') return next()
    if (!token) {
        next('/login')
    } else {
        next()
    }
})

export default router