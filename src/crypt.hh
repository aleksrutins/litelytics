#pragma once
#include <cstddef>
#include <stdexcept>
#include <string>

namespace litelytics::crypt {
    const unsigned char KEY_LENGTH_BYTES = 32;
    const unsigned char IV_LENGTH_BYTES = 16;
    using ustring = std::basic_string<std::byte>;
    ustring sha256(std::string);
    ustring aes(std::string);
}