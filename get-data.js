import * as db from './db/index.js';
import {log} from './log.js';
import express from 'express';
import CryptoJS from 'crypto-js';

/**
 * 
 * @param {express.Request} req 
 * @param {express.Response} res 
 */
export default async function getData(req, res) {
    log("User wants data");
    const siteData = await db.query(``, [req.params.site]);
    for(const site of userSites.rows) {
        siteData.push(await db.query(`
        `))
    }
    res.respondText(200, JSON.stringify({
        success: true,
        data: siteData.rows
    }));
}