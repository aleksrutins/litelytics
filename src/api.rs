use rocket::{Route, http::CookieJar};

#[get("/sites")]
pub async fn sites(
    cookies: &CookieJar<'_>
) {

}

pub fn api() -> Vec<Route> {
    routes![sites]
}