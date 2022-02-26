FROM node:alpine
ADD . /app
WORKDIR /app
RUN apk add curl
RUN yarn
ENTRYPOINT [ "yarn", "start" ]