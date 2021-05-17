const loginBtn         = document.querySelector('#login');
const createAccountBtn = document.querySelector('#create-account');
const email            = document.querySelector('#email');
const password         = document.querySelector('#password');
const tokenInput       = document.querySelector('#token');
const userIdInput      = document.querySelector('#user-id');
const loginFailed      = document.querySelector('#login-failed');
const createFailed     = document.querySelector('#create-failed');
async function login() {
    try {
        const result = await fetch(`/api/user/${email.value}/sign-in`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                password: password.value
            })
        });
        if(!result.ok) {
            throw new Error();
        }
        const json = await result.json();
        tokenInput.value = json.token;
        userIdInput.value = json.userId;
        window.opener.postMessage({msg: 'logged-in', detail: json}, "*");
    } catch(e) {
        loginFailed.removeAttribute('hidden');
    }
}
loginBtn.addEventListener('click', login);
createAccountBtn.addEventListener('click', async e => {
    try {
        const result = await fetch(`/api/user/${email.value}/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                password: password.value
            })
        });
        if(!result.ok) {
            throw new Error();
        }
        await login();
    } catch(e) {
        createFailed.removeAttribute('hidden');
    }
});