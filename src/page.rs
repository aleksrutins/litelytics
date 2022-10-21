macro_rules! basic_page {
    ($handler_name:ident, $url:expr, $template:expr, $title:expr) => {
        #[get($url)]
        async fn $handler_name() -> Template {
            Template::render(
                $template,
                context!(EmptyContext {
                    title: $title.into()
                }),
            )
        }
    };
}
macro_rules! dashboard_page {
    ($handler_name:ident, $url:expr, $template:expr) => {
        #[get($url)]
        async fn $handler_name(
            pool: &State<Pool<Postgres>>,
            cookies: &CookieJar<'_>,
        ) -> Result<Template, Redirect> {
            let Some(user_id) = cookies.get_private("user_id") else {
                return Err(Redirect::to(uri!(login)));
            };
        
            let email = if let Some(email) = cookies.get_private("user_email") {
                email.value().to_string()
            } else {
                sqlx::query!(
                    r"
        SELECT email FROM users
        WHERE id = $1
            ",
                    user_id
                        .value()
                        .parse::<i32>()
                        .or(Err(Redirect::to(uri!(login))))?
                )
                .fetch_one(pool.inner())
                .await
                .map_err(|_| Redirect::to(uri!(login)))?
                .email
            };
        
            Ok(Template::render(
                $template,
                context!(DashboardContext::new($url.to_string(), email)),
            ))
        }
    };
}