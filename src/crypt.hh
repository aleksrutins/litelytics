#pragma once
#include <cstddef>
#include <stdexcept>
#include <string>

namespace litelytics::crypt {
    using ustring = std::basic_string<std::byte>;
    ustring sha256(std::string);
}