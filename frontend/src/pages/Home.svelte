<script lang="ts">
    import { useSites } from '../api/sites';
import SitePreview from '../components/SitePreview.svelte';

    const sites = useSites()
</script>
<template>
    {#if $sites.data}
    <div class="p-3">
        <h2 class="text-xl font-bold">Your Sites</h2>
        <div class="flex flex-row flex-wrap">
            {#each $sites.data as site}
            <SitePreview {site}></SitePreview>
            {/each}
        </div>
    </div>
    {/if}
</template>