FROM debian

RUN apt-get update && apt-get install -y libpq-dev libpqxx-dev cmake ninja-build wget gcc 'g++' libasio-dev libboost-dev libboost-system-dev

RUN wget -O crow.deb "https://github.com/CrowCpp/Crow/releases/download/v1.0%2B5/crow-v1.0+5.deb"
RUN dpkg -i ./crow.deb

ADD . /app
WORKDIR /app

ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN mkdir build && cd build && cmake -GNinja .. && ninja

ARG SECRET_KEY
ENV SECRET_KEY=${SECRET_KEY}

CMD [ "./build/litelytics" ] 