FROM rust:alpine
ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

ADD . /app
WORKDIR /app

RUN cargo build --release

CMD [ "./target/litelytics" ]