console.log('Initializing Litelytics for ' + location.hostname)
const url = new URL(import.meta.url)
url.pathname = '/api/track'
fetch(url, {
    method: 'POST',
    body: JSON.stringify({
        'referrer': document.referrer
    }),
    headers: {
        'Content-Type': 'application/json'
    }
})