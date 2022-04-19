# Litelytics

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/THR6d_?referralCode=Y68pBw)

Litelytics is a lightweight, privacy-respecting analytics platform built using Node.js and PostgreSQL. It uses JWTs for secure authentication, and respects both Do Not Track and Global Privacy Control (by simply not logging anything).

## Deployment
I recommend that you click the button above. If, for some strange and incomprehensible reason, you would not like to host it on Railway, you can use [the Docker image](https://github.com/users/aleksrutins/packages/container/package/litelytics) (you need a Postgres database with `DATABASE_URL` for this to work):

```
docker pull ghcr.io/aleksrutins/litelytics:master
docker run -it ghcr.io/aleksrutins/litelytics
```

You can also host it from source (again, you need a Postgres database with `DATABASE_URL` set):
```
git clone https://github.com/aleksrutins/litelytics
cd litelytics
yarn
yarn start
```

You will also need to provide a `TOKEN_SECRET` environment to sign JWTs with. Preferably, generate it. (e.g. `openssl rand -hex 256`)

Then, run [`prepare-db.sql`](prepare-db.sql) to prepare the database.

## Usage
To use your instance, go to <https://litelytics-dashboard.vercel.app> (or [host it yourself](//github.com/aleksrutins/litelytics-dashboard); it's just a static site in `public/`). Enter your instance URL, and a login dialog will pop up. Enter your email address and preferred password, and click 'Create Account'. This will log you in. Then, add a site by inserting a row in the `sites` table and then hook that up with your user ID in the `usersites` table. THis will be automated eventually.

Then, add Litelytics to your site by adding the following to your `<head>`:
```html
<script type="module" src="https://your.litelytics.hostname/simpleclient.js"></script>
```
You will now see data in the dashboard (click 'View Data').