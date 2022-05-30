use sqlx::{postgres::PgPoolOptions, Pool, Postgres};
use std::env;

pub async fn get_pool() -> Pool<Postgres> {
    PgPoolOptions::new()
        .connect(&env::var("DATABASE_URL").expect("DATABASE_URL must be specified"))
        .await
        .expect("Failed to connect to database")
}
