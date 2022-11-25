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