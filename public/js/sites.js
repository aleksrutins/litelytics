const sites = await (await fetch("/api/sites")).json();
const app = document.querySelector("#app");

for(const site of sites) {
    console.log(site);
    const el = document.createElement('a');
    el.href = `/sites/${site.id}`;
    el.textContent = site.domain;
    el.classList.add("link");
    app.appendChild(el);
}