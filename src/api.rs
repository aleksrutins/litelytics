use rocket::{http::CookieJar, response::status::Unauthorized, Route, State};
use sqlx::{Pool, Postgres};

use crate::auth::ensure_authenticated;
use crate::models::{Site, UserSite, Visit, SiteData};

#[get("/sites")]
pub async fn sites(
    cookies: &CookieJar<'_>,
    pool: &State<Pool<Postgres>>,
) -> Result<String, Unauthorized<String>> {
    if !ensure_authenticated(cookies) {
        return Err(Unauthorized(Some("Not authenticated".into())));
    }
    let sites = sqlx::query_as!(
        Site,
        "select sites.* from usersites
        inner join sites on sites.id = site_id
        where user_id = $1",
        cookies
            .get_private("user_id")
            .map(|uid| uid
                .value()
                .parse::<i32>()
                .unwrap_or(-1)
            ).unwrap_or(-1)
    ).fetch_all(pool.inner())
    .await
    .map_err(|_| Unauthorized(Some("Database Error".to_string())))?;

    serde_json::to_string(&sites).map_err(|_| Unauthorized(Some("Error".to_string())))
}

#[post("/track")]
pub async fn track(pool: &State<Pool<Postgres>>, )

#[get("/sites/<id>")]
pub async fn site_data(id: i32, cookies: &CookieJar<'_>, pool: &State<Pool<Postgres>>) -> Result<String, Unauthorized<String>> {
    let Some(user_id) = cookies.get_private("user_id").map(|c| c.value().to_string()) else {
        return Err(Unauthorized(Some("Not authenticated".into())));
    };
    let usersite = sqlx::query_as!(
        UserSite,
        "select * from usersites
        where site_id = $1",
        id
    ).fetch_one(pool.inner()).await
    .map_err(|_| Unauthorized(Some("Database Error".to_string())))?;

    if usersite.user_id != user_id.parse::<i32>().expect("Failed to parse user ID") {
        return Err(Unauthorized(Some("Not authorized".into())));
    }

    let site = sqlx::query_as!(
        Site,
        "select * from sites
        where id = $1",
        id
    ).fetch_one(pool.inner()).await
    .map_err(|_| Unauthorized(Some("Database Error".to_string())))?;

    let visits = sqlx::query_as!(
        Visit,
        "select * from visits
        where site = $1",
        id
    ).fetch_all(pool.inner()).await
    .map_err(|_| Unauthorized(Some("Database Error".to_string())))?;

    serde_json::to_string(&SiteData {
        site, visits
    }).map_err(|_| Unauthorized(Some("Error".to_string())))
}
pub fn api() -> Vec<Route> {
    routes![sites, site_data]
}
