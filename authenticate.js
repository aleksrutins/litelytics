import jwt from 'jsonwebtoken';
import CryptoJS from 'crypto-js';
import * as db from './db/index.js';

export function generateAccessToken(userid, email) {
    return jwt.sign({
        id: userid,
        email
    }, process.env.TOKEN_SECRET, {expiresIn: "30m"});
}

export function checkToken(req, res, next) {
    const authHeader = req.headers['authorization'];
    const token = authHeader && authHeader.split(' ')[1];

    if (token == null) return res.respondText(403, 'Invalid token');

    jwt.verify(token, process.env.TOKEN_SECRET, (err, user) => {
        console.log(err);

        if (err) return res.respondText(403, "Invalid token");

        req.user = user;

        next();
    });
}

export async function signIn(req, res) {
    const password = req.body.password, email = req.params.email;
    const userDb = await db.query('SELECT * FROM users WHERE email = $1', [email]);
    if(!(CryptoJS.SHA256(password).toString() == userDb.rows[0].password)) {
        res.respondText(403, 'Incorrect password');
        return;
    }
    res.respondText(200, JSON.stringify({
        success: true,
        token: generateAccessToken(userDb.rows[0].id, email),
        userId: userDb.rows[0].id
    }));
}

export async function isAuthorizedForSite(user, site) {
    const rowCount = (await db.query('select * from usersites where user_id = $1 and site_id = $2', [user, site])).rowCount;
    if(rowCount == 0) {
        return false;
    }
    return true;
}