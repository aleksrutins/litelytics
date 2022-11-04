#include <cstdlib>
#include <cstring>
// adapted from https://github.com/devtomio/is-railway/blob/main/packages/node/src/index.ts
bool isRailway() {
    return getenv("RAILWAY_STATIC_URL")
        || getenv("RAILWAY_ENVIRONMENT")
        || getenv("RAILWAY_HEALTHCHECK_TIMEOUT_SEC")
        || getenv("RAILWAY_GIT_COMMIT_SHA")
        || getenv("RAILWAY_GIT_AUTHOR")
        || getenv("RAILWAY_GIT_BRANCH")
        || getenv("RAILWAY_GIT_REPO_NAME")
        || getenv("RAILWAY_GIT_REPO_OWNER")
        || getenv("RAILWAY_GIT_COMMIT_MESSAGE");
}
bool isProduction() {
    return getenv("RAILWAY_ENVIRONMENT") && !strcmp(getenv("RAILWAY_ENVIRONMENT"), "production");
}