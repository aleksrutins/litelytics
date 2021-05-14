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