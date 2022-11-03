import { LitElement, html, css } from "../lit-core.min.js";

class SiteLink extends LitElement {
    static properties = {
        id: {},
        domain: {}
    }

    static styles = css`
    a {
        padding: 7px;
        text-decoration: none;
        border: 1px solid black;
        color: black;
        border-radius: 4px;
        transition-property: color border-color;
        transition-duration: .2s
    }
    a:hover {
        color: var(--color-primary);
        border-color: var(--color-primary);
    }
    `

    render() {
        return html`
        <a href="/sites/${this.id}">${this.domain}</a>
        `
    }
}
customElements.define('site-link', SiteLink);