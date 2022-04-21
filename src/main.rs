mod nav;
#[macro_use] mod page_context;
mod schema;
mod models;
mod db;

#[macro_use] extern crate rocket;

#[macro_use]
extern crate diesel;
extern crate dotenv;

use diesel::{pg::PgConnection, Connection};
use dotenv::dotenv;
use nav::Nav;
use page_context::PageContext;
use rocket::fs::{FileServer, relative};
use rocket_dyn_templates::Template;
use std::{collections::HashMap, env};

#[get("/")]
fn index() -> Template {
    Template::render("index", context!(PageContext {
        nav: Nav::default(),
        url: "/".to_string(),
        title: "Home".to_string()
    }))
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
fn rocket() -> _ {
    let conn = db::establish_connection();

    rocket::build()
        .mount("/", routes![index, about])
        .mount("/public", FileServer::from(relative!("public")))
        .attach(Template::fairing())
}