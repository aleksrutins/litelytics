import { LitElement, html } from './lit-core.min.js';
import './components/loader.js';
import './components/site-link.js';

class AppView extends LitElement {
    static properties = {
        sites: {}
    }
    constructor() {
        super();
        this.sites = null;
        fetch('/api/sites')
            .then(res => res.json())
            .then(res => this.sites = res);
    }
    render() {
        return html`
            <h1>Your Sites</h1>
            ${this.sites? this.sites.map(site => html`
                <site-link id=${site.id} domain=${site.domain}></site-link>
            `) : html`<ll-loader></ll-loader>`}
        `
    }
}
customElements.define('ll-app', AppView);