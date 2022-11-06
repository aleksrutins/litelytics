#pragma once
#include <crow.h>
#include "app.hh"

namespace litelytics::auth::routes {
    void mount(litelytics::app::App *app);
}