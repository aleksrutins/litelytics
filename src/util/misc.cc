#include <cstdlib>
#include <fstream>

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
    optional<string> getAppOption(const char *env, const char *file) {
        auto value = getenv_opt(env);
        if(!value.has_value()) {
            ifstream fp(".pginfo");
            string url;
            getline(fp, url);
            value = url;
        }
        return value;
    }
}