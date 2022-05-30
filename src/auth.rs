#[derive(FromForm)]
pub struct LoginRequest<'r> {
    pub email: &'r str,
    pub password: &'r str,
}
