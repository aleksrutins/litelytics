<script setup lang="ts">
import Sidebar from './components/Sidebar.vue';
import { isAuthenticated } from './api/auth';
</script>

<template>
  <div class="flex flex-row w-screen h-screen">
    <Sidebar v-if="isAuthenticated()"></Sidebar>
    <div>
      <RouterView/>
    </div>
  </div>
</template>