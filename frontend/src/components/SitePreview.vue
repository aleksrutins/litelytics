<script setup lang="ts">
    import { SiteInfo } from '../api/sites';
import GeneratedIcon from './GeneratedIcon.vue';

    defineProps<{
        site: SiteInfo
    }>()
</script>
<template>
    <RouterLink class="shadow-lg hover:shadow-2xl dark:ring-1 dark:ring-slate-800 dark:hover:ring-2 transition rounded-lg p-3 m-2 flex flex-col items-center" :to="`/sites/${site.id}`">
        <img v-if="site.favicon" class="block w-[32px]" :src="site.favicon">
        <GeneratedIcon v-else size="32px"></GeneratedIcon>
        <p>{{site.domain}}</p>
    </RouterLink>
</template>