FROM rust:buster
ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

ADD . /app
WORKDIR /app

RUN apt-get install libpq-dev

RUN cargo build --release

CMD [ "./target/litelytics" ] 