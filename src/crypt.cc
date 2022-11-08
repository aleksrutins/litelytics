#include <openssl/evp.h>
#include <openssl/sha.h>
#include <openssl/rand.h>
#include <openssl/err.h>
#include <string>

#include "crypt.hh"
#include "util/except.hh"

using namespace std;
using namespace litelytics;

extern unsigned char *ll_secret_key;

namespace litelytics::crypt {
    char cryptErrSha256[] = "Error calculating SHA256 hash";
    using Sha256Exception = util::except::Exception<cryptErrSha256>;
    char cryptErrAES[] = "Error encrypting using AES";
    using AESException = util::except::Exception<cryptErrAES>;
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

    AESResult aes(string plaintext) {
        EVP_CIPHER_CTX *ctx;
        int len;
        int ciphertext_len;

        AESResult result;
        result.err = 0;

        result.ciphertext.reserve(((plaintext.length() / 8) + 2) * 8);

        if(!RAND_bytes(result.iv, IV_LENGTH_BYTES)) {
            result.err = ERR_get_error();
            return result;
        }

        if(!(ctx = EVP_CIPHER_CTX_new()))
            throw AESException();
        
        /*
        * Initialise the encryption operation. IMPORTANT - ensure you use a key
        * and IV size appropriate for your cipher
        * In this example we are using 256 bit AES (i.e. a 256 bit key). The
        * IV size for *most* modes is the same as the block size. For AES this
        * is 128 bits
        */
        if(1 != EVP_EncryptInit_ex(ctx, EVP_aes_256_cbc(), NULL, ll_secret_key, result.iv))
            throw AESException();

        /*
        * Provide the message to be encrypted, and obtain the encrypted output.
        * EVP_EncryptUpdate can be called multiple times if necessary
        */
        if(1 != EVP_EncryptUpdate(ctx, (unsigned char *) result.ciphertext.data(), &len, (const unsigned char *) plaintext.c_str(), plaintext.length()))
            throw AESException();
        ciphertext_len = len;

        /*
        * Finalise the encryption. Further ciphertext bytes may be written at
        * this stage.
        */
        if(1 != EVP_EncryptFinal_ex(ctx, ((unsigned char *) result.ciphertext.data()) + len, &len))
            throw AESException();
        ciphertext_len += len;

        /* Clean up */
        EVP_CIPHER_CTX_free(ctx);

        return result;
    }
}