#pragma once
#include <cstddef>
#include <stdexcept>
#include <string>

namespace litelytics::crypto {
    const unsigned char KEY_LENGTH_BYTES = 32;
    const unsigned char IV_LENGTH_BYTES = 16;
    using ustring = std::basic_string<std::byte>;
    ustring sha256(std::string);
    struct AESResult {
        ustring ciphertext;
        unsigned char iv[IV_LENGTH_BYTES];
        unsigned long err;
    };
    AESResult aes(std::string);
}