use serde::Serialize;

#[derive(Serialize)]
pub struct NavItem {
    pub label: String,
    pub url: String
}

#[derive(Serialize)]
pub struct Nav {
    pub items: Vec<NavItem>
}

macro_rules! nav_decl {
    ($($label:expr => $url:expr),*) => {
        Nav {
            items: vec![
                $(
                    NavItem {label: ($label).to_string(), url: ($url).to_string()}
                ),*
            ]
        }
    };
}

impl Nav {
    pub fn default() -> Nav {
        nav_decl!(
            "Home" => "/",
            "About" => "/about"
        )
    }
}