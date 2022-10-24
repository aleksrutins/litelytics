use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Site {
    pub id: i32,
    pub domain: String
}
