table! {
    sites (id) {
        id -> Int4,
        domain -> Text,
    }
}

table! {
    users (id) {
        id -> Int4,
        email -> Text,
        password -> Text,
    }
}

table! {
    usersites (id) {
        id -> Int4,
        user_id -> Int4,
        site_id -> Int4,
    }
}

table! {
    visits (id) {
        id -> Int4,
        site -> Int4,
        path -> Text,
        referer -> Text,
        timestamp -> Text,
        ip -> Nullable<Text>,
    }
}

joinable!(usersites -> sites (site_id));
joinable!(usersites -> users (user_id));
joinable!(visits -> sites (site));

allow_tables_to_appear_in_same_query!(
    sites,
    users,
    usersites,
    visits,
);
