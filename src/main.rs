mod auth;
#[macro_use]
mod page;
mod dashboard_context;
mod db;
mod nav;
mod api;
#[macro_use]
mod page_context;
mod models;

#[macro_use]
extern crate rocket;

extern crate dotenv;
extern crate sqlx;
use anyhow::Result;
use auth::LoginRequest;
use dashboard_context::DashboardContext;
use dotenv::dotenv;
use page_context::EmptyContext;
use rocket::{
    form::{Form, Strict},
    fs::{relative, FileServer},
    http::{Cookie, CookieJar},
    response::{
        status::{NotFound, Unauthorized},
        Redirect,
    },
    State,
};
use rocket_dyn_templates::Template;
use sqlx::{Pool, Postgres};
use std::env;

// Dashboard pages

dashboard_page!(index, "/", "index");

dashboard_page!(about, "/about", "about");

dashboard_page!(site_info, "/sites/<_>", "site_info");

// End dashboard pages

basic_page!(login, "/login", "login", "Log In");

#[post("/login", data = "<credentials>")]
async fn authenticate(
    pool: &State<Pool<Postgres>>,
    cookies: &CookieJar<'_>,
    credentials: Form<Strict<LoginRequest<'_>>>,
) -> Result<Redirect, NotFound<String>> {
    let req = credentials.into_inner();

    if let Ok(true) = auth::user_exists(&req.email, pool.inner()).await {
        if let Ok(id) = auth::authenticate(&req, pool.inner()).await {
            cookies.add_private(Cookie::new("user_id", id.to_string()));
            cookies.add_private(Cookie::new("user_email", req.email.to_string()));

            return Ok(Redirect::to(uri!(index)));
        }
    }

    return Err(NotFound(format!("Authentication failed")));
}

#[post("/create-account", data = "<credentials>")]
async fn create_account(
    pool: &State<Pool<Postgres>>,
    cookies: &CookieJar<'_>,
    credentials: Form<LoginRequest<'_>>,
) -> Result<Redirect, Unauthorized<String>> {
    let req = credentials.into_inner();

    let user_id = auth::create_account(&req, pool.inner()).await?;

    cookies.add_private(Cookie::new("user_id", user_id.to_string()));
    cookies.add_private(Cookie::new("user_email", req.email.to_string()));

    Ok(Redirect::to(uri!(index)))
}

#[get("/logout")]
fn logout(cookies: &CookieJar<'_>) -> Redirect {
    cookies.remove_private(Cookie::named("user_id"));
    cookies.remove_private(Cookie::named("user_email"));
    Redirect::to(uri!(login))
}

#[launch]
async fn rocket() -> _ {
    dotenv().ok();
    let pool = db::get_pool().await;
    rocket::build()
        .mount(
            "/",
            routes![index, about, login, logout, authenticate, create_account, site_info],
        )
        .mount("/api", api::api())
        .mount("/public", FileServer::from(relative!("public")))
        .manage(pool)
        .attach(Template::fairing())
}
