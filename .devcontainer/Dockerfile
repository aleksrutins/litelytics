FROM quay.io/fedora/fedora:37

RUN dnf install -y libpq-devel libpqxx-devel cmake ninja-build wget gcc 'gcc-c++' asio-devel boost-devel openssl-devel

RUN wget -O crow.tar.gz "https://github.com/CrowCpp/Crow/releases/download/v1.0%2B5/crow-v1.0+5.tar.gz"
RUN tar  -C /usr/local -xzf crow.tar.gz

COPY library-scripts/common-redhat.sh /tmp/library-scripts/
RUN bash /tmp/library-scripts/common-redhat.sh