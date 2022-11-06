#pragma once
#include <stdexcept>

namespace litelytics::util::except {
    template<char const *const message>
    class Exception : public std::exception {
    public:
        constexpr char *what() { 
            return message;
        }
    };
}