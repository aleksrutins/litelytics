#include <crow.h>
#include <crow/middlewares/session.h>
#include <crow/middlewares/cookie_parser.h>
#include <pqxx/pqxx>
#include <cstdlib>
#include <iostream>
#include <stdexcept>
#include <memory>
#include <openssl/aes.h>

#include "env.hh"
#include "app.hh"
#include "crypt.hh"
#include "auth/AuthController.hh"
#include "util/misc.hh"

using namespace litelytics;
using namespace litelytics::app;
using namespace litelytics::util;
using namespace std;

std::unique_ptr<pqxx::connection> ll_db_conn = nullptr;
unsigned char ll_secret_key[crypt::KEY_LENGTH_BYTES];
unsigned char ll_crypt_iv[crypt::IV_LENGTH_BYTES];

int main() {
    try {
        auto dburl = getAppOption("DATABASE_URL", ".pginfo");
        if(!dburl.has_value()) {
            cerr << "Error: Please provide either the DATABASE_URL environment variable or a database connection string in a .pginfo file, pointing to a valid PostgreSQL server." << endl;
            return 1;
        }

        auto secretKeyOpt = getAppOption("SECRET_KEY", ".secret-key");
        if(!secretKeyOpt.has_value()) {
            std::cerr << "Error: Please provide an encryption key in either the SECRET_KEY environmeny variable or a .secret-key file." << endl;
            return 1;
        }
        memcpy(ll_secret_key, secretKeyOpt.value().data(), crypt::KEY_LENGTH_BYTES);
        
        ll_db_conn = std::make_unique<pqxx::connection>(dburl.value());
        std::cout << "Connected to database" << std::endl;
        App app{Session{
            crow::FileStore{"/tmp/litelytics.session_data"}
        }};
        CROW_ROUTE (app, "/")([](crow::response &res) {
            res.set_static_file_info("static/index.html");
            res.end();
        });

        litelytics::auth::routes::mount(&app);

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
        return 0;
    } catch(std::exception const &e) {
        std::cerr << e.what() << std::endl;
        return 1;
    }
}
