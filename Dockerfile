FROM quay.io/fedora/fedora:37

RUN dnf install -y libpq-devel libpqxx-devel cmake ninja-build wget gcc 'gcc-c++' asio-devel boost-devel openssl-devel

RUN wget -O crow.tar.gz "https://github.com/CrowCpp/Crow/releases/download/v1.0%2B5/crow-v1.0+5.deb" && \
    tar -xzf crow.tar.gz -C /usr/local

ADD . /app
WORKDIR /app

ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN mkdir build && cd build && cmake -GNinja .. && ninja

ARG SECRET_KEY
ENV SECRET_KEY=${SECRET_KEY}

CMD [ "./build/litelytics" ] 