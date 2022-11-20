FROM quay.io/fedora/fedora:37

RUN dnf install -y git libpq-devel libpqxx-devel meson cmake ninja-build wget gcc 'gcc-c++' asio-devel boost-devel openssl-devel

RUN mkdir -p /opt
RUN git clone https://github.com/CrowCpp/Crow.git /opt/crow
WORKDIR /opt/crow
RUN mkdir build
WORKDIR /opt/crow/build
RUN cmake .. -DCROW_BUILD_EXAMPLES=OFF -DCROW_BUILD_TESTS=OFF
RUN make install

RUN dnf install -y gdb

COPY library-scripts/common-redhat.sh /tmp/library-scripts/
RUN bash /tmp/library-scripts/common-redhat.sh