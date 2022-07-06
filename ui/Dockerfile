FROM nginx:1.23-alpine
RUN mkdir -p /srv
COPY ./dist /srv/html
COPY ./config/nginx.dev.conf /etc/nginx/conf.d/default.conf
