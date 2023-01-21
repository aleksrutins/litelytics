<script lang="ts">
    import { active } from "tinro"
    import { getEmail, isAuthenticated, authState } from "../api/auth"
    import logo from '../assets/logo.webp'
    import * as cookie from '../util/cookie'

    function logOut() {
        cookie.del('userId')
        cookie.del('userEmail')
        authState.set(false)
    }
</script>
<div class="border-b z-10 bg-white/50 dark:bg-gray-900/50 backdrop-blur-md dark:border-slate-800 flex flex-row justify-between items-center font-stylized sticky top-0">
    <div class="flex flex-row items-center">
        <img src={logo} alt="Litelytics logo" width="50" class="self-center dark:invert"/>
        <a href="/" class="nav-link" use:active exact>Home</a>
    </div>
    <div class="pr-3">
        {#if isAuthenticated()}
        <button class="cursor-pointer" on:click={logOut}>Log Out {getEmail()}</button>
        {:else}
        Not signed in
        {/if}
    </div>
</div>