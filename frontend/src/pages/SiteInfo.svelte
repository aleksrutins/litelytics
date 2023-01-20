<script lang="ts">
    import { SiteInfo, useSite } from '../api/sites';
    import GeneratedIcon from '../components/GeneratedIcon.svelte';
    export let siteId: number
    let site = useSite(siteId);
</script>
<div class="p-3 max-h-screen flex">
    {#if $site.data}
    
    <div class="w-full flex items-center flex-col">
        {#if $site.data.site.favicon}
        <img src={$site.data.site.favicon} class="rounded-full block" width="100" alt={$site.data.site.domain}/>
        {:else}
        <GeneratedIcon size="100px" />
        {/if}
        <h1 class="text-xl pt-2">{$site.data.site.domain}</h1>
        <table class="shadow rounded-lg table-auto border-separate border overflow-auto block border-spacing-0 border-slate-300">
            <thead class="sticky top-0 bg-slate-200 border-slate-200 border-b shadow shadow-slate-100">
                <tr>
                    <th>Timestamp</th>
                    <th>Path</th>
                    <th>Referer</th>
                    <th>Client IP</th>
                </tr>
            </thead>
            <tbody>
                {#each $site.data.visits as visit}
                <tr>
                    <td><time datetime={visit.timestamp}>{new Date(visit.timestamp).toLocaleString()}</time></td>
                    <td>{visit.path}</td>
                    <td>{visit.referrer}</td>
                    <td>{visit.ip}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
    {:else}
    <div class="w-full flex items-center flex-col animate-pulse">
        <div class="rounded-full block w-[100px] h-[100px] bg-slate-200 dark:bg-slate-800"></div>
        <div class="h-7 mt-2 w-60 rounded bg-slate-200 dark:bg-slate-800"></div>
    </div>
    {/if}
</div>

<style>
thead, thead tr:first-of-type {
    border-top-left-radius: 00.5rem;
    border-top-right-radius: 00.5rem;
}
td, th {
    padding: 6px;
}

th {
    @apply text-slate-700;
}

tbody tr:not(:last-of-type) td {
    @apply border-b border-slate-100;
}
/* https://stackoverflow.com/a/47318412 */
th:first-of-type {
  border-top-left-radius: 0.5rem;
}
th:last-of-type {
  border-top-right-radius: 0.5rem;
}
tr:last-of-type td:first-of-type {
  border-bottom-left-radius: 0.5rem;
}
tr:last-of-type td:last-of-type {
  border-bottom-right-radius: 0.5rem;
}
</style>