FROM registry.redhat.io/ubi8/go-toolset as builder

RUN mkdir -p /opt/app-root/src
WORKDIR /opt/app-root/src

COPY go.* main.go /opt/app-root/src/

RUN go build -o hello-world

FROM registry.redhat.io/ubi8-minimal

USER root
RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /opt/app-root/src/hello-world /app/hello-world
USER 1001
EXPOSE 8080

CMD [ "./hello-world" ]
