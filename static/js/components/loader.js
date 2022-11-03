import { LitElement, css, html } from "../lit-core.min.js";

class Loader extends LitElement {
    static styles = css`
    @keyframes rotate {
        from {
            transform: rotate(0deg)
        }
        to {
            transform: rotate(360deg)
        }
    }    
    :host {
        border: 5px solid orange;
        border-radius: 30%;
        animation: rotate 1s linear 0s infinite;
        display: block;
        margin: 20px;
        width: 20px;
        height: 20px;
    }
    `
    render() {
        return html``;
    }
}
customElements.define('ll-loader', Loader);