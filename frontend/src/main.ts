import './style.css'
import App from './App.svelte'
import Home from './pages/Home.svelte'
import Login from './pages/Login.svelte'
import SiteInfo from './pages/SiteInfo.svelte'
import { initPathRouter, navigate } from '@bjornlu/svelte-router'
import { isAuthenticated } from './api/auth'

const router = initPathRouter([
    {
        path: '/',
        component: Home
    },
    {
        path: '/login',
        component: Login
    },
    {
        path: '/sites/:id',
        component: SiteInfo
    }
]);

const app = new App({
    target: document.body
})
if(!isAuthenticated()) {
    navigate("/login")
}