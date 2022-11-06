#include "mvc.hh"

namespace litelytics::util::mvc {
    void Controller::connect(crow::Crow<> &app) {
        auto routes = this->routes();
        for(RouteInfo route : routes) {
            app.route<crow::black_magic::get_parameter_tag(route.url)>(route.url)(route.handler);
        }
    }
}