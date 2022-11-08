litelytics_src = files(
    'main.cc', 'env.cc', 'db.cc', 'crypt.cc'
)

litelytics_deps = [
    dependency('Crow', method : 'cmake', modules : ['Crow::Crow']),
    dependency('libcrypto'),
    dependency('libpq'),
    dependency('libpqxx'),
]

subdir('auth')
subdir('util')

executable('litelytics', litelytics_src, dependencies : litelytics_deps)