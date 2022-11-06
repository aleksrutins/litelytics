#pragma once
#include <string>
#include <optional>
namespace litelytics::util {
    std::optional<std::string> getenv_opt(const char *);
}