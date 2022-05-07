use rocket::serde::Serialize;

use crate::nav::Nav;

#[derive(Serialize)]
pub struct PageContext {
    pub title: String,
    pub nav: Nav,
    pub url: String
}

#[derive(Serialize)]
pub struct EmptyContext {
    pub title: String
}

#[derive(Serialize)]
pub struct PageContextWrapper<T: Serialize> {
    pub page: T
}

macro_rules! context {
    ($ctx:expr) => {
        crate::page_context::PageContextWrapper {page: ($ctx)}
    };
}