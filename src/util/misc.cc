#include <cstdlib>

#include "misc.hh"

using namespace std;

namespace litelytics::util {
    optional<string> getenv_opt(const char *name) {
        auto env = getenv(name);
        if(env == nullptr) {
            return nullopt;
        } else {
            return make_optional<string>(env);
        }
    }
}