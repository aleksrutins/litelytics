import { isAuthorizedForSite } from './authenticate.js';
import * as db from './db/index.js';
export async function listSites(req, res) {
    const userSites = await db.query(`
select sites.id from usersites
    left join sites on sites.id = site_id
    where user_id = $1
    `, [req.user]);
    res.respondText(200, JSON.stringify({
        success: true,
        sites: userSites.rows.map(row => row.id)
    }));
}

export async function siteInfo(req, res) {
    if(!(await isAuthorizedForSite(req.user.id, req.params.site))) {
        res.respondText(403, JSON.stringify({
            success: false,
            err: 'EACCESS',
            detail: 'Unauthorized'
        }));
    }

    const siteInfo = await db.query('select * from sites where id = $1', [req.params.site]);
    res.respondText(200, JSON.stringify({
        success: true,
        data: siteInfo.rows[0]
    }));
}