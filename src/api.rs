use rocket::{http::CookieJar, response::status::Unauthorized, Route, State};
use sqlx::{Pool, Postgres};

use crate::auth::ensure_authenticated;
use crate::models::Site;

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
    .map_err(|_| Unauthorized(Some("Could not fetch sites".to_string())))?;

    serde_json::to_string(&sites).map_err(|_| Unauthorized(Some("Error".to_string())))
}

pub fn api() -> Vec<Route> {
    routes![sites]
}
