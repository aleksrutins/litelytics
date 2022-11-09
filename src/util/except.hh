#pragma once
#include <stdexcept>

namespace litelytics::util::except {
    template<char const *const message>
    class Exception : public std::exception {
    public:
        virtual const char *what() { 
            return message;
        }
    };
}