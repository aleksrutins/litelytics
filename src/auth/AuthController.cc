#include "AuthController.hh"
#include "util.hh"
#include "db.hh"

using namespace litelytics;
using namespace litelytics::app;
using namespace litelytics::auth::util;

namespace litelytics::auth::routes {
    App *_globalApp;

    bool authenticate(crow::request req, char *email, char *password) {
        if(!checkCredentials(email, password)) {
            return false;
        }
        pqxx::work txn{db::conn()};
        int userId = txn.query_value<int>("SELECT id FROM users WHERE email = " + txn.esc(email));
        txn.commit();
        return true;
    }

    auto login(crow::request req) {
        if(req.method == "GET"_method) {
            auto page = crow::mustache::load("login.html");
            return page.render();
        } else if(req.method == "POST"_method) {
            
        }
    }

    void create_account(crow::request req, crow::response &res) {
        res.redirect("/");
        res.end();
    }

    void mount(App *app) {
        _globalApp = app;
        CROW_ROUTE((*app), "/auth/login").methods("GET"_method, "POST"_method)(login);
        CROW_ROUTE((*app), "/auth/create-account").methods("POST"_method)(create_account);
    }
}