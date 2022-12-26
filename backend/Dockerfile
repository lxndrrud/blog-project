FROM node:16.13.2-alpine

RUN apk update && apk add tzdata
ENV TZ=Europe/Moscow

COPY . /usr/src/app

WORKDIR /usr/src/app

USER node

ENTRYPOINT [ "npm", "start" ]

EXPOSE 8001