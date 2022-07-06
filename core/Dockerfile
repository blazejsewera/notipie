# make build-docker before building docker image
FROM scratch
COPY ./notipie /
COPY ./notipie.config.json /
EXPOSE 8080
STOPSIGNAL SIGINT
ENTRYPOINT ["/notipie"]
