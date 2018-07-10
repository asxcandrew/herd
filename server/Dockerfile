FROM alpine:3.7
ENV GIN_MODE release
WORKDIR /home/builds
COPY build /home/builds
COPY config.toml /home/builds
ENTRYPOINT ["./server"]
EXPOSE 8080
