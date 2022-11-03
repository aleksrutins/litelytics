#include <crow.h>

int main() {
    crow::SimpleApp app;
    CROW_ROUTE (app, "/")([](crow::response &res) {
        res.set_static_file_info("static/index.html");
        res.end();
    });
    app.port(8080).multithreaded().run();
    return 0;
}