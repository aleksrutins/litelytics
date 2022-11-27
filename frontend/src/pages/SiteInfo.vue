<script setup lang="ts">
    import { useRoute } from 'vue-router';
    import { useSite } from '../api/sites';
    import GeneratedIcon from '../components/GeneratedIcon.vue';

    const route = useRoute()
    const site = useSite(parseInt(route.params.id as string))
</script>
<template>
    <div class="p-3">
        <div v-if="site.data.value" class="w-full flex items-center flex-col">
            <img v-if="site.data.value?.site.favicon" :src="site.data.value.site.favicon" class="rounded-full block" width="100"/>
            <GeneratedIcon size="100px" v-else/>
            <h1 class="text-xl pt-2">{{site.data.value?.site.domain}}</h1>
        </div>
        <div v-else class="w-full flex items-center flex-col animate-pulse">
            <div class="rounded-full block w-[100px] h-[100px] bg-slate-200 dark:bg-slate-800"></div>
            <div class="h-7 mt-2 w-60 rounded bg-slate-200 dark:bg-slate-800"></div>
        </div>
    </div>
</template>