use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Site {
    pub id: i32,
    pub domain: String,
}

#[derive(Serialize, Deserialize)]
pub struct UserSite {
    pub id: i32,
    pub site_id: i32,
    pub user_id: i32,
}

#[derive(Serialize, Deserialize)]
pub struct Visit {
    pub id: i32,
    pub site: i32,
    pub path: String,
    pub referer: String,
    pub timestamp: String,
    pub ip: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub struct SiteData {
    pub site: Site,
    pub visits: Vec<Visit>,
}
