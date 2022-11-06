FROM quay.io/fedora/fedora:37

RUN dnf install -y libpq-devel libpqxx-devel cmake ninja-build wget gcc git 'gcc-c++' asio-devel boost-devel openssl-devel

RUN mkdir -p /opt
RUN git clone https://github.com/CrowCpp/Crow.git /opt/crow
WORKDIR /opt/crow
RUN mkdir build
WORKDIR /opt/crow/build
RUN cmake .. -DCROW_BUILD_EXAMPLES=OFF -DCROW_BUILD_TESTS=OFF
RUN make install

ADD . /app
WORKDIR /app

ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN mkdir build && cd build && cmake -GNinja .. && ninja

ARG SECRET_KEY
ENV SECRET_KEY=${SECRET_KEY}

ARG PORT
ENV PORT=${PORT}

CMD [ "./build/litelytics" ] 