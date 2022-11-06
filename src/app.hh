#pragma once
#include <crow.h>
#include <crow/middlewares/session.h>
#include <crow/middlewares/cookie_parser.h>

namespace litelytics::app {
    using Session = crow::SessionMiddleware<crow::FileStore>;
    using App = crow::App<crow::CookieParser, Session>;
}