<script setup lang="ts">
    import { useRoute } from 'vue-router';
    import { useSite } from '../api/sites';
    import GeneratedIcon from '../components/GeneratedIcon.vue';

    const route = useRoute()
    const site = useSite(parseInt(route.params.id as string))
</script>
<template>
    <div class="p-3 max-h-screen flex">
        <div v-if="site.data.value" class="w-full flex items-center flex-col">
            <img v-if="site.data.value?.site.favicon" :src="site.data.value.site.favicon" class="rounded-full block" width="100"/>
            <GeneratedIcon size="100px" v-else/>
            <h1 class="text-xl pt-2">{{site.data.value?.site.domain}}</h1>
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
                    <tr v-for="visit in site.data.value.visits">
                        <td><time :datetime="visit.timestamp">{{new Date(visit.timestamp).toLocaleString()}}</time></td>
                        <td>{{visit.path}}</td>
                        <td>{{visit.referrer}}</td>
                        <td>{{visit.ip}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div v-else class="w-full flex items-center flex-col animate-pulse">
            <div class="rounded-full block w-[100px] h-[100px] bg-slate-200 dark:bg-slate-800"></div>
            <div class="h-7 mt-2 w-60 rounded bg-slate-200 dark:bg-slate-800"></div>
        </div>
    </div>
</template>

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