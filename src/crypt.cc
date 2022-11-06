#include <openssl/evp.h>
#include <openssl/sha.h>
#include <string>

#include "crypt.hh"
#include "util/except.hh"

using namespace std;
using namespace litelytics;

namespace litelytics::crypt {
    char cryptErrSha256[] = "Error calculating SHA256 hash";
    using Sha256Exception = util::except::Exception<cryptErrSha256>;
    // https://stackoverflow.com/a/2262447
    ustring sha256(string input) {
        EVP_MD_CTX *context = EVP_MD_CTX_new();

        size_t len = input.length();
        ustring out;
        out.reserve(SHA256_DIGEST_LENGTH);

        if(!EVP_DigestInit_ex(context, EVP_get_digestbyname("SHA256"), NULL))
            throw Sha256Exception();

        if(!EVP_DigestUpdate(context, input.c_str(), len))
            throw Sha256Exception();

        if(!EVP_DigestFinal(context, (unsigned char *) out.data(), NULL))
            throw Sha256Exception();

        return out;
    }
}