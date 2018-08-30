FROM alpine:3.8
ENV GIN_MODE release
WORKDIR /home/builds
COPY build /home/builds
ENTRYPOINT ["./server"]
EXPOSE 8080
