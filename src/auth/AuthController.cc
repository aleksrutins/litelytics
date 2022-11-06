#include "AuthController.hh"
#include "util.hh"
#include "db.h"

using namespace litelytics;
using namespace litelytics::app;
using namespace litelytics::auth::util;

namespace litelytics::auth::routes {
    App &_globalApp;

    bool authenticate(crow::request req, char *email, char *password) {
        if(!checkCredentials(email, password)) {
            return false;
        }
        ppqx::work txn{db::conn()};
        int userId = txn.query_value<int>("SELECT id FROM users WHERE email = " + txn.esc(email));
        txn.commit();
        
    }

    auto login(crow::request req) {
        if(req.method == "GET"_method) {
            auto page = crow::mustache::load("login.html");
            return page.render();
        } else if(req.method == "POST"_method) {
            
        }
    }

    auto create_account(crow::request req) {

    }

    void mount(App &app) {
        _globalApp = app;
        CROW_ROUTE(app, "/auth/login").methods("GET"_method, "POST"_method)(login);
        CROW_ROUTE(app, "/auth/create-account").methods("POST"_method)(create_account);
    }
}