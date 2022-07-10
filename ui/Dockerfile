# build
FROM node:16-bullseye AS build-env

ARG NODE_ENV=production
ENV NODE_ENV $NODE_ENV
WORKDIR /ui

# preload dependencies
COPY .yarnrc.yml package.json yarn.lock ./
COPY .yarn ./.yarn
RUN yarn install

COPY . .

RUN yarn build
RUN mv dist /
RUN mv config/nginx.dev.conf /

# run
FROM nginx:1.23

RUN mkdir -p /srv

COPY --from=build-env /dist /srv/html
COPY --from=build-env /nginx.dev.conf /etc/nginx/conf.d/default.conf
