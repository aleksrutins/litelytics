#include <string>

#include "auth.hh"
#include "../db.hh"
#include "../crypt.hh"
using namespace litelytics;
using namespace std;
namespace litelytics::auth::util {
    bool checkCredentials(string email, string passwd) {
        if(!db::isConnected()) return false;
        try {
            bool result = false;
            pqxx::work txn{db::conn()};
            auto expectedHash = crypt::string_to_ustring(txn.query_value<string>(
                "SELECT password FROM users"
                "WHERE email = '" + txn.esc(email) + "'"
            ));
            crypt::ustring actualHash = crypt::sha256(passwd);
            if(actualHash == expectedHash) result = true;
            txn.commit();
            return result;
        } catch(const std::exception &e) {
            return false;
        }
    }
}
