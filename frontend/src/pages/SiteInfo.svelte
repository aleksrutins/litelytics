<script lang="ts">
    import { SiteInfo, useSite } from "../api/sites";
    import VisitsChart from "../components/data-display/VisitsChart.svelte";
    import VisitsTable from "../components/data-display/VisitsTable.svelte";
    import GeneratedIcon from "../components/GeneratedIcon.svelte";
    export let siteId: number;
    let site = useSite(siteId);

    let displays = new Map<string, [ConstructorOfATypedSvelteComponent, string][]>([
        ["Table", [
            [VisitsTable, "Visits Table"]
        ]],
        ["Chart", [
            [VisitsChart, "Visits Per Week"]
        ]]
    ]);
    let currentDisplayName = "Table";
    $: currentDisplay = displays.get(currentDisplayName)!;

    function setDisplay(ev: MouseEvent) {
        currentDisplayName = (<Node>ev.target).textContent!
    }
</script>

<div class="h-full flex">
    {#if $site.data}
        <div class="w-full h-full flex flex-row items-center">
            <div class="flex flex-col p-2 border-r dark:border-gray-800 m-0 h-full">
                <div class="flex flex-shrink flex-row">
                    {#if $site.data.site.favicon}
                        <img
                            src={$site.data.site.favicon}
                            class="rounded-full block align-middle"
                            width="32"
                            alt={$site.data.site.domain}
                        />
                    {:else}
                        <GeneratedIcon size="32px" />
                    {/if}
                    <h1 class="font-stylized font-bold pl-2">
                        {$site.data.site.domain}
                    </h1>
                </div>
                {#each Array.from(displays.keys()) as name}
                    <button class="block m-2 p-2 rounded text-start" class:bg-gray-100={currentDisplayName == name} class:dark:bg-gray-700={currentDisplayName == name} on:click={setDisplay}>{name}</button>
                {/each}
            </div>
            <div class="flex flex-row items-start flex-wrap flex-grow p-3 overflow-auto h-full">
                {#each currentDisplay as component}
                    <div class="resize border dark:border-gray-800 rounded-md flex flex-col">
                        <div class="border-b dark:border-gray-800 bg-gray-100 dark:bg-gray-900 rounded-t-md px-1 text-sm font-bold">
                            {component[1]}
                        </div>
                        <div class="p-2 w-[20rem] h-[20rem] flex">
                            <svelte:component this={component[0]} visits={$site.data.visits} />
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    {:else}
        <div class="w-full flex flex-row animate-pulse">
            <div
                class="rounded-full block w-[32px] h-[32px] bg-slate-200 dark:bg-slate-800"
            />
            <div class="h-7 ml-2 w-60 rounded bg-slate-200 dark:bg-slate-800" />
        </div>
    {/if}
</div>
