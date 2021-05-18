import * as db from './db/index.js';
import express from 'express';
import { isAuthorizedForSite } from './authenticate.js';

/**
 * 
 * @param {express.Request} req 
 * @param {express.Response} res 
 */
export default async function getData(req, res) {
    if(!(await isAuthorizedForSite(req.user.id, req.params.site))) {
        res.respondText(403, JSON.stringify({
            success: false,
            err: 'EACCESS',
            detail: 'Not authorized'
        }));
        return;
    }
    const siteData = (await db.query(`select * from visits where site = $1`, [req.params.site])).rows;
    res.respondText(200, JSON.stringify({
        success: true,
        data: siteData.rows
    }));
}