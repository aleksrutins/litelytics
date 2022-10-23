use std::error::Error;

use rocket::{http::CookieJar, response::status::Unauthorized, Route, State};
use sqlx::{Pool, Postgres};

use crate::auth::ensure_authenticated;

#[get("/sites")]
pub async fn sites(
    cookies: &CookieJar<'_>,
    pool: &State<Pool<Postgres>>,
) -> Result<String, Unauthorized<String>> {
    if !ensure_authenticated(cookies) {
        return Err(Unauthorized(Some("Not authenticated".into())));
    }
    let sites = sqlx::query!(
        "select sites.* from usersites
        left join sites on sites.id = site_id
        where user_id = $1",
        cookies
            .get("user_id")
            .unwrap()
            .value()
            .parse::<i32>()
            .unwrap()
    );
    Ok("".into())
}

pub fn api() -> Vec<Route> {
    routes![sites]
}
