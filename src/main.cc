#include <crow.h>
#include <pqxx/pqxx>
#include <cstdlib>
#include <iostream>
#include <stdexcept>

pqxx::connection *dbconn = nullptr;

int main() {
    try {
        auto dburl = getenv("DATABASE_URL");
        if(dburl == nullptr) {
            std::cerr << "Error: Please provide the DATABASE_URL environment variable, pointing to a valid PostgreSQL server." << std::endl;
            return 1;
        }
        dbconn = new pqxx::connection(dburl);
        std::cout << "Connected to database" << std::endl;
        crow::SimpleApp app;
        CROW_ROUTE (app, "/")([](crow::response &res) {
            res.set_static_file_info("static/index.html");
            res.end();
        });
        auto port_str = getenv("PORT");
        int port = 
            port_str == nullptr
                ? 8080
                : atoi(port_str);
        app.port(port).multithreaded().run();
        return 0;
    } catch(std::exception const &e) {
        std::cerr << e.what() << std::endl;
        return 1;
    }
}
