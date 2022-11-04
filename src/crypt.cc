#include <openssl/sha.h>

namespace litelytics::crypt {
    // https://stackoverflow.com/a/2262447
    bool sha256(void *input, size_t length, unsigned char *out) {
        SHA256_CTX context;

        if(!SHA256_Init(&context))
            return false;

        if(!SHA256_Update(&context, input, length))
            return false;

        if(!SHA256_Final(out, &context))
            return false;

        return true;
    }
}