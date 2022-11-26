<script setup lang="ts">
    import { useSites } from '../api/sites';
import SitePreview from '../components/SitePreview.vue';

    const sites = useSites()
</script>
<template>
    <div class="p-3" v-if="sites.data.value">
        <h2 class="text-xl font-bold">Your Sites</h2>
        <div class="flex flex-row flex-wrap">
            <SitePreview v-for="site in sites.data.value" :site="site"></SitePreview>
        </div>
    </div>
</template>