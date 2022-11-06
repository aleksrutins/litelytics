#pragma once
#include <crow.h>
#include "app.hh"

namespace litelytics::auth {
    class AuthController {
    public:
        static void mount(litelytics::app::App &app);
    };
}