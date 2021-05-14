import * as db from './db/index.js';
import { log } from './log.js';
export default async function track(req, res) {
    log("Getting client");
    const client = await db.getClient();
    try {
        const siteId = await (await client.query('select id from sites where domain = $1', [data.site])).rows[0].id;
        log("begin");
        await client.query('BEGIN');
        const query = 'INSERT INTO visits(site, useragent, referrer, path, timestamp) VALUES($1, $2, $3, $4, $5) RETURNING id';
        log("insert");
        let data = req.body;
        console.log("Headers", req.headers);
        data.useragent = req.headers['user-agent'];
        if('referrer' in req.headers) data.referrer = req.headers['referer'];
        const result = await db.query(query, [siteId, data.useragent, data.referrer || 'UNKNOWN', data.path, new Date().toString()]);
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
            detail: e.stack
        }));
    } finally {
        client.release();
    }
}