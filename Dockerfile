FROM rust:buster
ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN apt-get install libpq-dev

RUN rustup default nightly

ADD . /app
WORKDIR /app

RUN cargo build --release

CMD [ "./target/release/litelytics" ] 