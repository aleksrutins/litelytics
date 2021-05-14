FROM node:alpine
ADD . /workspace
WORKDIR /workspace
RUN apk add curl
RUN curl -f https://get.pnpm.io/v6.js | node - add --global pnpm
RUN pnpm i
CMD [ "pnpm", "start" ]