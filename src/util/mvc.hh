#include <vector>
#include <functional>
#include <crow/app.h>
namespace litelytics::util::mvc {
    struct RouteInfo {
        char const *const url;
        void *handler;
    };
    class Controller {
    protected:
        virtual std::vector<RouteInfo> routes() = 0;
    public:
        void connect(crow::Crow<> &app);
    };
}