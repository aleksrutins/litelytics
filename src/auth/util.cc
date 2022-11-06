#include <string>

#include "util.hh"
#include "../db.hh"
#include "../crypt.hh"
using namespace litelytics;
using namespace std;
namespace litelytics::auth::util {
    bool checkCredentials(char *email, char *passwd) {
        if(!db::isConnected()) return false;
        try {
            bool result = false;
            pqxx::work txn{db::conn()};
            auto expectedHash = txn.query_value<crypt::ustring>(
                "SELECT password FROM users"
                "WHERE email = '" + txn.esc(email) + "'"
            );
            auto actualHash = crypt::sha256(passwd);
            if(actualHash == expectedHash) result = true;
            txn.commit();
            return result;
        } catch(const std::exception &e) {
            return false;
        }
    }
}
