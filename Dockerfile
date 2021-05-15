FROM node:alpine
ADD . /workspace
WORKDIR /workspace
RUN apk add curl
RUN yarn
CMD [ "yarn", "start" ]