<script>
    export let visits
</script>
<table class="shadow rounded-lg table-auto border-separate border overflow-auto block border-spacing-0 border-slate-300 dark:border-gray-800">
    <thead class="sticky top-0 bg-slate-200 border-slate-200 border-b shadow shadow-slate-100 dark:border-gray-800">
        <tr>
            <th>Timestamp</th>
            <th>Path</th>
            <th>Referer</th>
            <th>Client IP</th>
        </tr>
    </thead>
    <tbody>
        {#each visits as visit}
        <tr>
            <td><time datetime={visit.timestamp}>{new Date(visit.timestamp).toLocaleString()}</time></td>
            <td>{visit.path}</td>
            <td>{visit.referrer}</td>
            <td>{visit.ip}</td>
        </tr>
        {/each}
    </tbody>
</table>

<style>
    thead,
    thead tr:first-of-type {
        border-top-left-radius: 00.5rem;
        border-top-right-radius: 00.5rem;
    }
    td,
    th {
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