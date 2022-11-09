litelytics_src = files(
    'main.cc', 'env.cc', 'db.cc', 'crypto.cc'
)

litelytics_deps = [
    dependency('Crow', method : 'cmake'),
    dependency('libcrypto'),
    dependency('libpq'),
    dependency('libpqxx'),
]

subdir('auth')
subdir('util')

executable('litelytics', litelytics_src, dependencies : litelytics_deps, include_directories : include_directories('.'))