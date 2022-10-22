FROM rust:slim AS build

RUN apk add libpq-dev

RUN rustup default nightly

ADD . /app
WORKDIR /app

ARG DATABASE_URL
ENV DATABASE_URL=${DATABASE_URL}

RUN cargo build --release

FROM rust:slim
ARG ROCKET_ADDRESS=0.0.0.0
ARG ROCKET_SECRET_KEY
ENV ROCKET_ADDRESS=${ROCKET_ADDRESS}
ENV ROCKET_SECRET_KEY=${ROCKET_SECRET_KEY}

COPY --from=build /app/target/release/litelytics .
COPY --from=build /app/public .
COPY --from=build /app/templates .

CMD [ "./litelytics" ] 