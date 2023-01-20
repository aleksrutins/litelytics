<script lang="ts">
    import { SiteInfo } from '../api/sites';
import GeneratedIcon from './GeneratedIcon.svelte';

    export let site: SiteInfo
</script>

<a class="ring-1 ring-slate-200 dark:ring-slate-800 hover:ring-2 transition rounded-lg p-3 m-2 flex flex-col items-center" href={`/sites/${site.id}`}>
    {#if site.favicon}
    <img class="block w-[32px]" src={site.favicon} alt={site.domain}>
    {:else}
    <GeneratedIcon size="32px"></GeneratedIcon>
    {/if}
    <p>{site.domain}</p>
</a>