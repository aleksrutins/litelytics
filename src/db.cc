#include <pqxx/pqxx>
#include "db.hh"

extern pqxx::connection *ll_db_conn;

namespace litelytics::db {
    bool isConnected() {
        return ll_db_conn != nullptr;
    }
    pqxx::connection *ref() {
        return ll_db_conn;
    }
}