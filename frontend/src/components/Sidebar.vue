<script setup lang="ts">
    import { useSites } from '../api/sites';
    import { authState, getEmail } from '../api/auth';
    import * as cookie from '../util/cookie';
    import { useRouter } from 'vue-router';

    const sites = useSites()
    const linkStyle = "rounded hover:bg-emerald-500 px-3 py-2 transition"
    const router = useRouter()
    function logOut() {
        cookie.del('userId')
        cookie.del('userEmail')
        authState.loggedIn = false
        router.push('/login')
    }
</script>
<template>
    <div class="flex flex-shrink flex-col items-stretch justify-between bg-primary-light p-3">
        <div class="flex flex-col">
            <img src="../assets/logo.webp" width="100" class="self-center"/>
            <RouterLink :class="linkStyle" to="/">Home</RouterLink>
            <h2 class="font-bold uppercase text-xs text-gray-800 mx-3 pt-2">Your sites</h2>
            <RouterLink v-for="site of sites.data.value" :to="'/sites/' + site.id" :class="linkStyle">
                <img v-if="site.favicon != null" :src="site.favicon" width="24" class="rounded-full inline align-middle pr-1"/>
                <span class="align-middle">{{site.domain}}</span>
            </RouterLink>
        </div>
        <div class="flex flex-col">
            <span>{{getEmail()}}</span>
            <a class="cursor-pointer" @click="logOut">Log Out</a>
        </div>
    </div>
</template>