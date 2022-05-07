#![feature(let_else)]
mod nav;
mod db;
mod auth;
mod dashboard_context;
#[macro_use] mod page_context;

#[macro_use] extern crate rocket;

extern crate sqlx;
extern crate dotenv;
use auth::LoginRequest;
use dashboard_context::DashboardContext;
use dotenv::dotenv;
use nav::Nav;
use page_context::{PageContext, EmptyContext};
use rocket::{fs::{FileServer, relative}, State, response::{Redirect, status::{NotFound, Unauthorized}}, http::{CookieJar, Cookie}, form::{Form, Strict}};
use rocket_dyn_templates::Template;
use sha2::{Sha256, Digest};
use sqlx::{Pool, Postgres, Executor};
use std::{env, any::Any};

fn hash(str: String) -> String {
    let mut hasher = Sha256::new();
    hasher.update(str.as_bytes());
    format!("{:X}", hasher.finalize())
}

#[get("/login")]
async fn login() -> Template {
    Template::render("login", context!(EmptyContext{title: "Log In".into()}))
}

#[post("/login", data = "<credentials>")]
async fn authenticate(pool: &State<Pool<Postgres>>, cookies: &CookieJar<'_>, credentials: Form<Strict<LoginRequest<'_>>>) -> Result<Redirect, NotFound<String>> {
    let req = credentials.into_inner();
    let record = sqlx::query!(r"
SELECT * FROM users
WHERE email = $1
    ", req.email).fetch_one(pool.inner()).await.or(Err(NotFound("User not found".to_string())))?;

    let pass_hash = hash(req.password.to_string());
 
    if pass_hash != record.password {
        return Err(NotFound("User not found".into()))
    }
    
    cookies.add_private(Cookie::new("user_id", record.id.to_string()));

    Ok(Redirect::to(uri!(index)))
}

#[post("/create-account", data = "<credentials>")]
async fn create_account(pool: &State<Pool<Postgres>>, cookies: &CookieJar<'_>, credentials: Form<LoginRequest<'_>>) -> Result<Redirect, Unauthorized<String>> {
    let req = credentials.into_inner();
    let exists = sqlx::query!(r"
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
)
    ", req.email).fetch_one(pool.inner()).await.or(Err(Unauthorized(Some("Could not fetch user existence state".to_string()))))?.exists;
    if exists != Some(false) {
        return Err(Unauthorized(Some("User already exists".to_string())));
    }
    let password_hash = hash(req.password.to_string());
    let user_id = sqlx::query!(r"
INSERT INTO users ( email, password )
VALUES ( $1, $2 )
RETURNING id
    ", req.email, password_hash).fetch_one(pool.inner()).await.or(Err(Unauthorized(Some("Could not create account".to_string()))))?.id;
    
    cookies.add_private(Cookie::new("user_id", user_id.to_string()));

    Ok(Redirect::to(uri!(index)))
}

#[get("/")]
async fn index(pool: &State<Pool<Postgres>>, cookies: &CookieJar<'_>) -> Result<Template, Redirect> {
    let Some(user_id) = cookies.get_private("user_id") else {
        return Err(Redirect::to(uri!(login)));
    };

    let email = sqlx::query!(r"
SELECT email FROM users
WHERE id = $1
    ", user_id.value().parse::<i32>().or(Err(Redirect::to(uri!(login))))?).fetch_one(pool.inner()).await.or(Err(Redirect::to(uri!(login))))?.email;

    Ok(Template::render("index", context!(DashboardContext::new("/".to_string(), email))))
}

#[get("/about")]
fn about() -> Template {
    Template::render("about", context!(PageContext {
        nav: Nav::default(),
        url: "/about".to_string(),
        title: "About".to_string()
    }))
}

#[launch]
async fn rocket() -> _ {
    dotenv().ok();
    let pool = db::get_pool().await;
    rocket::build()
        .mount("/", routes![
            index,
            about,
            login,
            authenticate,
            create_account
            ])
        .mount("/public", FileServer::from(relative!("public")))
        .manage(pool)
        .attach(Template::fairing())
}
