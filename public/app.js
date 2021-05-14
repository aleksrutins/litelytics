const logOutBtn = document.querySelector('#logOutBtn');
const siteName = document.querySelector('#site-name');
logOutBtn.addEventListener('click', e => {
    localStorage.removeItem('token');
    localStorage.removeItem('email');
    location.assign('/login.html');
});

console.log("Getting data...");
(async () => {
    const result = await fetch('/sites/', {
        method: 'POST',
        headers: {
            'Authorization': 'Bearer ' + localStorage['token']
        }
    });
})();