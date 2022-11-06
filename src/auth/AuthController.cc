#include "AuthController.hh"

using namespace litelytics::app;

namespace litelytics::auth {
    auto login(crow::request &req) {
        auto page = crow::mustache::load("login.html");
        return page.render();
    }

    void AuthController::mount(App &app) {
        CROW_ROUTE(app, "/auth/login")(login);
    }
}