mod auth;
mod dashboard_context;
mod db;
mod nav;
#[macro_use]
mod page_context;

#[macro_use]
extern crate rocket;

extern crate dotenv;
extern crate sqlx;
use anyhow::Result;
use auth::LoginRequest;
use dashboard_context::DashboardContext;
use dotenv::dotenv;
use nav::Nav;
use page_context::{EmptyContext, PageContext};
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

#[get("/login")]
async fn login() -> Template {
    Template::render(
        "login",
        context!(EmptyContext {
            title: "Log In".into()
        }),
    )
}

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

#[get("/")]
async fn index(
    pool: &State<Pool<Postgres>>,
    cookies: &CookieJar<'_>,
) -> Result<Template, Redirect> {
    let Some(user_id) = cookies.get_private("user_id") else {
        return Err(Redirect::to(uri!(login)));
    };

    let email = if let Some(email) = cookies.get_private("user_email") {
        email.value().to_string()
    } else {
        sqlx::query!(
            r"
SELECT email FROM users
WHERE id = $1
    ",
            user_id
                .value()
                .parse::<i32>()
                .or(Err(Redirect::to(uri!(login))))?
        )
        .fetch_one(pool.inner())
        .await
        .map_err(|_| Redirect::to(uri!(login)))?
        .email
    };

    Ok(Template::render(
        "index",
        context!(DashboardContext::new("/".to_string(), email)),
    ))
}

#[get("/about")]
fn about() -> Template {
    Template::render(
        "about",
        context!(PageContext {
            nav: Nav::default(),
            url: "/about".to_string(),
            title: "About".to_string()
        }),
    )
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
            routes![index, about, login, logout, authenticate, create_account],
        )
        .mount("/public", FileServer::from(relative!("public")))
        .manage(pool)
        .attach(Template::fairing())
}
