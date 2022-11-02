use rocket::{
    http::{uri::Authority, Status},
    outcome::Outcome,
    request::{self, FromRequest},
    State,
};
use serde::{Deserialize, Serialize};
use sqlx::{Pool, Postgres};

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
    pub timestamp: i64,
    pub ip: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateVisit {
    pub site: i32,
    pub path: String,
    pub referer: String,
    pub timestamp: i64,
    pub ip: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateVisitMeta {
    pub site: i32,
    pub timestamp: i64,
    pub ip: Option<String>
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateVisitRequest {
    pub path: String,
    pub referer: String
}

#[derive(Serialize, Deserialize)]
pub struct SiteData {
    pub site: Site,
    pub visits: Vec<Visit>,
}

#[rocket::async_trait]
impl<'r> FromRequest<'r> for CreateVisitMeta {
    type Error = ();

    async fn from_request(request: &'r rocket::Request<'_>) -> request::Outcome<Self, Self::Error> {
        let pool = match <&State<Pool<Postgres>>>::from_request(request).await {
            Outcome::Success(pool) => pool,
            Outcome::Failure(e) => return Outcome::Failure(e),
            Outcome::Forward(f) => return Outcome::Forward(f),
        };
        let domain = Authority::parse(
            request
                .headers()
                .get_one("Origin")
                .expect("Origin header not found"),
        )
        .unwrap();
        let site = if let Ok(site) = sqlx::query!(
            "select id from sites
            where domain = $1",
            domain.host()
        )
        .fetch_one(pool.inner())
        .await
        .map(|record| record.id)
        {
            site
        } else {
            return Outcome::Failure((Status::InternalServerError, ()));
        };
        Outcome::Success(
            Self {
                site,
                ip: if Some("1") == request.headers().get_one("Sec-GPC") || Some("1") == request.headers().get_one("DNT") {
                    None
                } else {
                    request.client_ip().map(|ip| ip.to_string())
                },
                timestamp: chrono::offset::Utc::now().timestamp()
            }
        )
    }
}
