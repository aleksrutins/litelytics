import { LitElement, html } from "./lit-core.min.js";
import './components/loader.js';

class AppView extends LitElement {
    static properties = {
        data: {},
        error: {}
    }
    constructor() {
        super();
        this.data = null;
        this.error = false;
        const id = location.pathname.split('/').pop();
        fetch(`/api/sites/${id}`)
            .then(res => res.json())
            .then(res => this.data = res)
            .then(_ => console.log(this.data))
            .catch(e => {
                console.log(e);
                this.error = true;
            });
    }
    render() {
        return html`
            ${
            this.error
                ? html`<span class="error">Error fetching site data</span>`
                : this.data
                    ? html`
                        <h1>${this.data.domain}</h1>
                    `
                    : html`<ll-loader></ll-loader>`
            }
        `
    }
}
customElements.define('ll-app', AppView);