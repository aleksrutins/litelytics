#include <crow.h>
#include <pqxx/pqxx>
#include <cstdlib>
#include <iostream>
#include <stdexcept>

#include "env.hh"

pqxx::connection *ll_db_conn = nullptr;

void cleanup() {
    if(ll_db_conn) {
        ll_db_conn->close();
        delete ll_db_conn;
    }
}

int main() {
    try {
        auto dburl = getenv("DATABASE_URL");
        if(dburl == nullptr) {
            std::cerr << "Error: Please provide the DATABASE_URL environment variable, pointing to a valid PostgreSQL server." << std::endl;
            return 1;
        }
        ll_db_conn = new pqxx::connection(dburl);
        std::cout << "Connected to database" << std::endl;
        crow::SimpleApp app;
        CROW_ROUTE (app, "/")([](crow::response &res) {
            res.set_static_file_info("static/index.html");
            res.end();
        });
        auto portStr = getenv("PORT");
        int port = 
            portStr
                ? atoi(portStr)
                : 8080;
        if(isRailway()) {
            std::cout << "Running in Railway!" << std::endl;
            if(isProduction()) {
                std::cout << "Running in production; some logs will be silenced" << std::endl;
                app.loglevel(crow::LogLevel::WARNING);
            }
        }
        std::cout << "\e[1mStarting server on port " << port << "\e[0m" << std::endl;
        app.port(port).multithreaded().run();
        cleanup();
        return 0;
    } catch(std::exception const &e) {
        std::cerr << e.what() << std::endl;
        cleanup();
        return 1;
    }
}
