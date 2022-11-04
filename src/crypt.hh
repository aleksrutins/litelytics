#pragma once
#include <cstddef>
namespace litelytics::crypt {
    bool sha256(void *input, size_t length, unsigned char *out);
}