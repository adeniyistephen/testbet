FROM golang:alpine
ENV CGO_ENABLED 0
ARG VCS_REF

COPY . /testbet

WORKDIR /testbet

RUN go build -ldflags "-X main.build=${VCS_REF}"

CMD [ "./testbet" ]