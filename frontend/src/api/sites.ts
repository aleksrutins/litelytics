import { useQuery } from "@tanstack/vue-query"
import { getEmail } from "./auth"

export type SiteInfo = {
    id: number,
    domain: string,
    favicon?: string
}

export type Visit = {
    id: number,
    path: string,
    referrer: string,
    timestamp: string,
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
        queryKey: ['sites', { user: getEmail() }],
        queryFn: () => fetch('/api/sites').then(r => r.json())
    })
}

export function useSite(id: number) {
    return useQuery<SiteData>({
        queryKey: ['site', id, { user: getEmail() }],
        queryFn: () => fetch('/api/sites/' + id).then(r => r.json())
    })
}