import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Home from './pages/Home.vue'
import Login from './pages/Login.vue'
import SiteInfo from './pages/SiteInfo.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { isAuthenticated } from './api/auth'
import { VueQueryPlugin } from '@tanstack/vue-query'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Home },
        { path: '/login', component: Login },
        { path: '/sites/:id', component: SiteInfo }
    ]
})

const app = createApp(App)

app.use(router)
app.use(VueQueryPlugin)

app.mount("#app")

if(!isAuthenticated()) {
    router.push("/login")
}