export function all() {
    const cookiesStr = document.cookie.split(';').map(s => s.trim())
    const cookies = new Map<string, string>()
    for(const cookie of cookiesStr) {
        const parts = cookie.split('=')
        cookies.set(parts[0], parts[1])
    }
    return cookies
}
export function has(name: string) {
    return document.cookie.includes(name + '=')
}

export function get(name: string) {
    return all().get(name)
}

export function del(name: string) {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`
}