import * as cookie from '../util/cookie'

export function isAuthenticated() {
    return document.cookie.includes("userId=") && document.cookie.includes("userEmail=")
}

export function getEmail() {
    return cookie.get('userEmail')
}