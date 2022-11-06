#include <pqxx/pqxx>
#include <memory>
#include "db.hh"

using namespace std;

extern unique_ptr<pqxx::connection> ll_db_conn;

namespace litelytics::db {
    bool isConnected() {
        return ll_db_conn != nullptr;
    }
    pqxx::connection &conn() {
        return *ll_db_conn;
    }
}