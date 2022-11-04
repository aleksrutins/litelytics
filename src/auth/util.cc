#include <string>

#include "auth.hh"
#include "../db.hh"
#include "../crypt.hh"
using namespace litelytics;
using namespace std;
namespace litelytics::auth::util {
    bool checkCredentials(char *email, char *passwd) {
        if(!db::isConnected())
        try {
            pqxx::work txn{*db::conn()};
        } catch(const std::exception &e) {

        }
    }
}