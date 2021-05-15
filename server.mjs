import express from 'express';
import * as http from 'http';
import { Server } from 'socket.io';
import cors from 'cors';
import { log } from './log.js';
import { resolveRel } from './util.mjs';
import track from './track.js';
import { addSite, addUser, addUserToSite } from './add-site.js';
import getData from './get-data.js';
import { checkToken, signIn } from './authenticate.js';
import { listSites } from './list-sites.js';

const app = express();
const server = http.createServer(app);
const io = new Server(server);

app.use((req, res, next) => {
    log(`\x1B[36m${req.method}\x1B[0m \x1B[32m${req.path}\x1B[0m`);
    res.respond = (status, filename) => {
        log(`\x1b[36mRESPOND\x1b[0m \x1b[32m${status}\x1b[0m ${filename}`);
        res.status(status);
        res.sendFile(resolveRel(filename, import.meta));
    };
    res.respondText = (status, text) => {
        log(`\x1b[36mRESPOND\x1b[0m \x1b[32m${status}\x1b[0m TEXT`);
        res.status(status);
        res.send(text);
    };
    res.api = (fn) => {
        log(`\x1b[36mAPI\x1b[0m ${fn.name}`);
        return fn(req, res);
    }
    next();
});

app.use(express.static('public'));

app.options('/api/site/*/track', cors());
app.post('/api/site/:domain/track', express.json(), cors(), async (req, res) => {
    // validation
    const body = req.body;
    if(!('path' in body)) {
        log("Not enough columns");
        res.respondText(500, JSON.stringify({
            success: false,
            err: 'ENOTENOUGH',
            detail: 'Not enough columns in query; need path'
        }));
        return;
    }
    log("Beginning `track`");
    await res.api(track);
});

app.post('/api/site/:name/create', express.json(), async (req, res) => {
    // validation
    const body = req.body;
    if(!('password' in body)) {
        console.log(body);
        log("Not enough columns");
        res.respondText(500, JSON.stringify({
            success: false,
            err: 'ENOTENOUGH',
            detail: 'Not enough columns in query; need password'
        }));
        return;
    }
    await res.api(addSite);
});

app.get('/api/site/list', checkToken, listSites);
app.get('/api/site/:site/data', express.json(), checkToken, getData);

app.post('/api/user/:email/sign-in', express.json(), signIn);
app.post('/api/user/:email/create', express.json(), addUser);
app.post('/api/site/:siteId/user/:userId/add', checkToken, addUserToSite);


app.get('/log.js', (req, res) => {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.respond(200, 'log.js');
});

app.get('/simpleclient.js', (req, res) => {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.respond(200, 'simpleclient.js');
});

server.listen(process.env.PORT || 3000, () => {
    log("Listening on 3000");
});