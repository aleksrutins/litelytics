#include <string>

#include "AuthController.hh"
#include "util.hh"
#include "db.hh"
#include "crypt.hh"

using namespace std;
using namespace litelytics;
using namespace litelytics::app;
using namespace litelytics::auth::util;

namespace litelytics::auth::routes {
    App *_globalApp;

    bool authenticateSimple(crow::request req, int userId, string email) {
        auto encryptedId = crypt::aes(to_string(userId));
        if(encryptedId.err != 0) return false;
        auto encryptedEmail = crypt::aes(email);
        if(encryptedEmail.err != 0) return false;
        auto idStr = (encryptedId.ciphertext + (byte)'/') + (byte*)encryptedId.iv;
        auto emailStr = (encryptedEmail.ciphertext + (byte)'/') + (byte*)encryptedEmail.iv;
        
        auto &session = _globalApp->get_context<Session>(req);

        session.set("userId", crow::utility::base64encode((char *)idStr.c_str(), idStr.size()));
        session.set("userEmail", crow::utility::base64encode((char *)emailStr.c_str(), emailStr.size()));
        return true;
    }

    bool authenticate(crow::request req, char *email, char *password) {
        if(!checkCredentials(email, password)) {
            return false;
        }
        pqxx::work txn{db::conn()};
        int userId = txn.query_value<int>("SELECT id FROM users WHERE email = '" + txn.esc(email) + "'");
        txn.commit();

        return authenticateSimple(req, userId, email);
    }

    void login(crow::request req, crow::response &res) {
        auto page = crow::mustache::load("login.html");
        if(req.method == "GET"_method) {
            crow::mustache::context ctx;
            char *err;
            if((err = req.url_params.get("err")) != nullptr) {
                ctx["error"] = err;
            }
            res.end(page.render_string(ctx));
        } else if(req.method == "POST"_method) {
            crow::query_string data(req.body, false);
            if(authenticate(req, data.get("email"), data.get("password"))) {
                res.redirect("/");
                res.end();
            } else {
                crow::mustache::context ctx;
                ctx["error"] = "Invalid credentials";
                res.end(page.render_string(ctx));
            }
        }
    }

    void createAccount(crow::request req, crow::response &res) {
        crow::query_string data(req.body, false);
        char *email = data.get("email")
           , *password = data.get("password");
        try {
            pqxx::work txn{db::conn()};
            int userId = txn.query_value<int>("SELECT id FROM users WHERE email = '" + txn.esc(email) + "'");
            txn.commit();
            res.redirect("/login?err=User%20already%20exists");
            res.end();
            return;
        } catch(std::exception _e) {
            auto passwordHash = crypt::sha256(password);
            auto &conn = db::conn();
            conn.prepare("createUser",
                "INSERT INTO USERS (email, password)"
                "VALUES ($1, $2)"
                "RETURNING id"
            );
            pqxx::work txn{conn};
            auto result = txn.exec_prepared1("createUser", email, passwordHash);
            txn.commit();
            if(!authenticateSimple(req, result[0].as<int>(), email)) {
                res.redirect("/login?err=Error%20creating%20user");
                res.end();
                return;
            };
            res.redirect("/");
            res.end();
        }
    }

    void mount(App *app) {
        _globalApp = app;
        CROW_ROUTE((*app), "/auth/login").methods("GET"_method, "POST"_method)(login);
        CROW_ROUTE((*app), "/auth/create-account").methods("POST"_method)(createAccount);
    }
}