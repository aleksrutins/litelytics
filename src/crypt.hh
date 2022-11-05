#pragma once
#include <cstddef>
#include <stdexcept>
#include <string>

namespace litelytics::crypt {
    using ustring = std::basic_string<unsigned char>;
    ustring string_to_ustring(std::string str);
    ustring sha256(std::string);
}