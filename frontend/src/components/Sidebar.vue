<script setup lang="ts">
    import { useSites } from '../api/sites';
    import { authState, getEmail } from '../api/auth';
    import * as cookie from '../util/cookie';
    import { useRouter } from 'vue-router';

    const sites = useSites()
    const linkStyle = "rounded hover:bg-emerald-500 mb-[1px] dark:hover:bg-emerald-800 px-3 py-2 transition"
    const activeLinkStyle = "bg-primary-dark dark:bg-emerald-900"
    const sectionHeaderStyle = "font-bold uppercase text-xs text-gray-800 dark:text-slate-200 mx-3 pt-2"
    const router = useRouter()
    function logOut() {
        cookie.del('userId')
        cookie.del('userEmail')
        authState.loggedIn = false
        router.push('/login')
    }
</script>
<template>
    <div class="flex flex-shrink flex-col items-stretch justify-between bg-primary-light dark:bg-slate-900 dark:text-white p-3">
        <div class="flex flex-col">
            <img src="../assets/logo.webp" width="100" class="self-center dark:invert"/>
            <RouterLink :class="linkStyle + ' ' + (router.currentRoute.value.path == '/' ? activeLinkStyle : '')" to="/">Home</RouterLink>
            <h2 :class="sectionHeaderStyle">Your Sites</h2>
            <RouterLink v-for="site of sites.data.value" :to="'/sites/' + site.id" :class="linkStyle + ' ' + (router.currentRoute.value.path == ('/sites/' + site.id.toString()) ? activeLinkStyle : '')">
                <img v-if="site.favicon != null" :src="site.favicon" width="24" class="rounded-full inline align-middle pr-1"/>
                <span class="align-middle">{{site.domain}}</span>
            </RouterLink>
        </div>
        <div class="flex flex-col">
            <h2 :class="sectionHeaderStyle">{{getEmail()}}</h2>
            <a class="cursor-pointer" :class="linkStyle" @click="logOut">Log Out</a>
        </div>
    </div>
</template>