FROM golang:1.16

RUN mkdir /notipie
ADD . /notipie/
WORKDIR /notipie

ENV PORT=5000
EXPOSE 5000

RUN make build
CMD ["./notipie"]
