use paste::paste;
use super::schema::*;
macro_rules! model {
    ($table:expr, $name:ident; $id_field:ident => $($fname:ident => $type:ty),*) => {
        #[derive(Queryable)]
        pub struct $name {
            pub $id_field: i32,
            $(
                pub $fname: $type
            ),*
        }
        paste! {
            #[derive(Insertable)]
            #[table_name=$table]
            pub struct [<New $name>] {
                $(
                    pub $fname: $type
                ),*
            }
        }
    };
}

model!("users", User; id =>
    email => String,
    password => String
);

model!("sites", Site; id =>
    domain => String
);

model!("usersites", UserSite; id =>
    user_id => i32,
    site_id => i32
);

model!("visits", Visit; id =>
    site => i32,
    path => String,
    referer => String,
    timestamp => String,
    ip => Option<String>
);