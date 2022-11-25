<script setup lang="ts">
import Sidebar from './components/Sidebar.vue';
import { authState } from './api/auth'; 
</script>

<template>
  <div class="flex flex-row w-screen h-screen bg-primary-light dark:bg-slate-900">
    <Sidebar v-if="authState.loggedIn"></Sidebar>
    <div class="flex-grow bg-white dark:bg-black dark:text-white rounded-tl-lg">
      <RouterView/>
    </div>
  </div>
</template>