use rocket::{request::{FromRequest, self}, http::uri::Uri};
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

#[derive(Debug, Serialize, Deserialize)]
pub struct Visit {
    pub id: i32,
    pub site: i32,
    pub path: String,
    pub referer: String,
    pub timestamp: String,
    pub ip: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateVisit {
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



impl<'a, 'r> FromRequest<'r> for CreateVisit {
    type Error = Box<dyn std::error::Error>;

    async fn from_request(request: &'r rocket::Request<'a>) -> request::Outcome<Self, Self::Error> {
        let domain = Uri::parse(request.headers().get_one("Origin").expect("Origin header not found"))?.authority().unwrap().host();
        let site = sqlx::query!(
            "select id from sites
            where domain = $1",
            domain
        );
        // rocket::outcome::Outcome::Success(
        //     Self {

        //     }
        // )
        todo!()
    }
}