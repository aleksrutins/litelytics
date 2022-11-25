<script setup lang="ts">
    import { useSites } from '../api/sites';
    import { getEmail } from '../api/auth';
    const sites = useSites()
    const linkStyle = "rounded hover:bg-emerald-500 px-3 py-2 transition"
</script>
<template>
    <div class="flex flex-shrink flex-col items-stretch justify-between bg-primary-light p-3">
        <div class="flex flex-col">
            <img src="../assets/logo.webp" width="100" class="self-center"/>
            <RouterLink :class="linkStyle" to="/">Home</RouterLink>
            <h2 class="font-bold uppercase text-xs text-gray-800 mx-3 pt-2">Your sites</h2>
            <RouterLink v-for="site of sites.data.value" :to="'/sites/' + site.id" :class="linkStyle">{{site.domain}}</RouterLink>
        </div>
        <div class="flex flex-col">
            <span>{{getEmail()}}</span>
            <a href="/auth/logout">Log Out</a>
        </div>
    </div>
</template>