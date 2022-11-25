import { useQuery } from "@tanstack/vue-query"

export type SiteInfo = {
    id: number,
    domain: string,
    favicon?: string
}

export type Visit = {
    id: number,
    path: string,
    referer: string,
    timestamp: Date,
    ip: String
}

export type SiteData = {
    id: number,
    site: {
        domain: string,
        favicon?: string
    },
    visits: Visit[]
}

export function useSites() {
    return useQuery<SiteInfo[]>({
        queryKey: ['sites'],
        queryFn: () => fetch('/api/sites').then(r => r.json())
    })
}

export function useSite(id: number) {
    return useQuery<SiteData>({
        queryKey: ['site', id],
        queryFn: () => fetch('/api/sites/' + id).then(r => r.json())
    })
}