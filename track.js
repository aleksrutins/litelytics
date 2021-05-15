import * as db from './db/index.js';
export default async function track(req, res) {
    let data = req.body;
    const client = await db.getClient();
    try {
        const siteId = await (await client.query('select id from sites where domain = $1', [req.params.domain])).rows[0].id;
        await client.query('BEGIN');
        
        const query = 'INSERT INTO visits(site, useragent, referer, path, timestamp, ip) VALUES($1, $2, $3, $4, $5, $6)';
        
        data.useragent = req.headers['user-agent'];
        if('referer' in req.headers) data.referer = req.headers['referer'];
        
        const result = await db.query(query, [siteId, data.useragent, data.referer || 'UNKNOWN', data.path, new Date().toString(), req.ip]);
        
        await client.query('COMMIT');
        
        res.respondText(200, JSON.stringify({
            success: true
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