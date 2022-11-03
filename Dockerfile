FROM debian AS build

RUN apt-get update && apt-get install -y libpq-dev meson wget

RUN wget -O crow.deb "https://github.com/CrowCpp/Crow/releases/download/v1.0%2B5/crow-v1.0+5.deb"
RUN dpkg -i ./crow.deb

ADD . /app
WORKDIR /app

ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN meson builddir
RUN ninja -C builddir

ARG SECRET_KEY
ENV SECRET_KEY=${SECRET_KEY}

CMD [ "./builddir/src/litelytics" ] 