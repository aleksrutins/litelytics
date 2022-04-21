use rocket::serde::{Serialize, Deserialize};

use crate::nav::Nav;

#[derive(Serialize)]
pub struct PageContext {
    pub title: String,
    pub nav: Nav,
    pub url: String
}

#[derive(Serialize)]
pub struct PageContextWrapper {
    pub page: PageContext
}

macro_rules! context {
    ($ctx:expr) => {
        crate::page_context::PageContextWrapper {page: ($ctx)}
    };
}