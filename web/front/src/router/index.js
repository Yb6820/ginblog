import Vue from 'vue'
import VueRouter from 'vue-router'
import ArticleList from '../components/ArticleList.vue'
import Detail from '../components/Details.vue'
import Home from "../views/Home.vue"

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'Home',
    component: Home,
    children: [
        { path: '/', component: ArticleList, meta: { title: '欢迎来到GinBlog' } },
        { path: '/detail/:id', component: Detail, meta: { title: '文章详情' }, props: true }
    ]
}, ]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
})

//导航守卫
router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})
export default router