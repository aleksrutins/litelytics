<script lang="ts">
    import { Visit } from "../../api/sites";
    import { VisXYContainer, VisLine, VisAxis } from "@unovis/svelte"


    export let visits: Visit[]

    function visitsToData(visits: Visit[]) {
        let result = new Map<number, number>()
        for(const visit of visits) {
            const day = Math.round(new Date(visit.timestamp).getTime() / (1000 * 60 * 60 * 24 * 7))
            if(result.has(day)) result.set(day, result.get(day)! + 1)
            else result.set(day, 1)
        }
        return Array.from(result.entries());
    }

    $: data = visitsToData(visits)

    const tickFormat = (tick: number) => new Date(tick * 1000 * 60 * 60 * 24 * 7).toDateString()
</script>

<VisXYContainer {data}>
    <VisLine curveType="step" {data} x={d => d[0]} y={d => d[1]} />
    <VisAxis type="x" {tickFormat}/>
    <VisAxis type="y"/>
</VisXYContainer>