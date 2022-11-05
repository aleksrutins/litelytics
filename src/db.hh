#pragma once
#include <pqxx/pqxx>

namespace litelytics::db {
    bool isConnected();
    pqxx::connection *ref();
}