import Analytics from 'https://cdn.skypack.dev/litelytics-client';
import { log } from './log.js';
log.nocolor = true;
log("Initializing Litelytics for " + location.hostname);
const url = new URL(import.meta.url);
(new Analytics(url.protocol + '//' + url.host)).track();