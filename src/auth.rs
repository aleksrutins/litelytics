use anyhow::Result;
use rocket::response::status::Unauthorized;
use sha2::{Digest, Sha256};
use sqlx::{Pool, Postgres};
use std::str;

#[derive(FromForm)]
pub struct LoginRequest<'r> {
    pub email: &'r str,
    pub password: &'r str,
}

pub async fn user_exists(email: &str, pool: &Pool<Postgres>) -> Result<bool> {
    sqlx::query!(
        r"
        SELECT EXISTS(
            SELECT * FROM users
            WHERE email = $1
        )
        ",
        email
    )
    .fetch_one(pool)
    .await
    .map_err(|_| anyhow::anyhow!("Could not fetch user existence state"))
    .map(|row| row.exists.unwrap_or(false))
}

pub async fn authenticate(req: &LoginRequest<'_>, pool: &Pool<Postgres>) -> Result<i32> {
    let pass_hash = hash(req.password.to_string());
    let record = sqlx::query!(
        r"
        SELECT * FROM users
        WHERE email = $1
        ",
        req.email
    )
    .fetch_one(pool)
    .await
    .map_err(|_| anyhow::anyhow!("Failed to fetch record"))?;

    if str::from_utf8(&record.password).unwrap() == pass_hash {
        Ok(record.id)
    } else {
        Err(anyhow::anyhow!("Failed to authenticate"))
    }
}

pub async fn create_account(
    req: &LoginRequest<'_>,
    pool: &Pool<Postgres>,
) -> Result<i32, Unauthorized<String>> {
    if user_exists(req.email, pool).await.unwrap_or(true) {
        return Err(Unauthorized(Some("User already exists".to_string())));
    }

    let password_hash = hash(req.password.to_string());
    let user_id = sqlx::query!(
        r"
        INSERT INTO users(email, password)
        VALUES ($1, $2)
        RETURNING id
        ",
        req.email,
        password_hash.as_bytes()
    )
    .fetch_one(pool)
    .await
    .map_err(|_| Unauthorized(Some("Failed to create user".to_string())))?
    .id;

    Ok(user_id)
}

fn hash(str: String) -> String {
    let mut hasher = Sha256::new();
    hasher.update(str.as_bytes());
    format!("{:X}", hasher.finalize())
}
