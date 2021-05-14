import { log } from '/log.js';

const loginBtn    = document.querySelector('#login--login-btn');
const createBtn   = document.querySelector('#login--create-btn');
const email       = document.querySelector('#login--email');
const siteName    = document.querySelector('#login--site');
const password    = document.querySelector('#login--password');
const errorBanner = document.querySelector('#errorBanner');
log.nocolor = true;

createBtn.addEventListener('click', async e => {
    try {
        log(`Creating user ${email.value}...`);
        await fetch(`/api/user/${email.value}/create`, {
            method: 'POST',
            body: JSON.stringify({
                password: password.value
            })
        })
        signIn();
    } catch(e) {
        log("There was an error.");
        errorBanner.textContent = e.message;
        errorBanner.style.display = 'block';
    }
});

loginBtn.addEventListener('click', async function signIn() {
    log('Signing in');
    try {
        let response = await fetch('/sign-in', {
            method: 'POST',
            body: JSON.stringify({
                email: email.value,
                password: password.value
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if(!response.ok) {
            log("There was an error.");
            throw new Error(await response.text());
        }
        const tokenJSON = await response.json();
        localStorage['token'] = tokenJSON.token;
        localStorage['site'] = siteName.value;
        location.assign('/');
    } catch(e) {
        log("There was an error.");
        errorBanner.textContent = e.message;
        errorBanner.style.display = 'block';
    }
});