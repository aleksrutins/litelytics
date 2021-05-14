import * as db from './db/index.js';
import CryptoJS from 'crypto-js';
import { log } from './log.js';
import { isAuthorizedForSite } from './authenticate.js';

export async function addUser(req, res) {
    const data = req.body;
    log("Getting client");
    const client = await db.getClient();
    try {
        log("Checking if user already exists");
        let exists = await (await client.query('SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)', [data.email])).rows[0].exists;
        if(exists) throw new Error('User exists');
        log("begin");
        await client.query('BEGIN');
        const query = 'INSERT INTO users(email, password) VALUES($1, $2) RETURNING id';
        log("insert");
        const result = await db.query(query, [data.email, CryptoJS.SHA256(data.password).toString()]);
        log('commit');
        await client.query('COMMIT');
        log('respond');
        res.respondText(200, JSON.stringify({
            success: true,
            id: result.rows[0].id
        }));
    } catch(e) {
        await client.query('ROLLBACK');
        res.respondText(500, JSON.stringify({
            success: false,
            err: 'EDATABASE',
            detail: e.message
        }));
    } finally {
        client.release();
    }
}

export async function addSite(req, res) {
    log('Getting client')
    const client = await db.getClient();
    try {
        let exists = await (await client.query('SELECT EXISTS(SELECT 1 FROM sites WHERE domain = $1)', [req.params.name])).rows[0].exists;
        if(exists) throw new Error('Site exists');
        
        await client.query('BEGIN');
        const result = await client.query('INSERT INTO sites(name) VALUES($1) RETURNING id', [req.params.name]);
        await client.query('COMMIT');
        await client.query('BEGIN');
        const addUserResult = await client.query('INSERT INTO usersites(user_id, site_id) VALUES((SELECT id FROM users WHERE email = $1), $2)', [req.user, result.rows[0].id]);
        await client.query('COMMIT');

        res.respondText(200, JSON.stringify({
            success: true,
            id: result.rows[0].id
        }));
    } catch(e) {
        await client.query('ROLLBACK');
        res.respondText(500, JSON.stringify({
            success: false,
            err: 'EDATABASE',
            detail: e.message
        }));
    } finally {
        client.release();
    }
}

export async function addUserToSite(req, res) {
    const client = await db.getClient();
    if(isAuthorizedForSite(req.user.id, req.params.site))
    try {
        
    } catch(e) {
        await client.query('ROLLBACK');
        res.respondText(500, JSON.stringify({
            success: false,
            err: 'EDATABASE',
            detail: e.message
        }));
    } finally {
        client.release();
    }
}