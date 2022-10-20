use serde::Serialize;

use crate::nav::Nav;

#[derive(Serialize)]
pub struct DashboardContext {
    pub title: String,
    pub nav: Nav,
    pub url: String,
    pub me: String,
}

impl DashboardContext {
    pub fn new(url: String, me: String) -> DashboardContext {
        DashboardContext {
            title: "Dashboard".to_string(),
            nav: Nav::default(),
            url,
            me,
        }
    }
}
