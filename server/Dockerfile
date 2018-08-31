FROM alpine:3.8
ENV GIN_MODE release
WORKDIR /home/builds
COPY build /home/builds
COPY access_model.conf /home/builds
COPY access_policy.csv /home/builds
ENTRYPOINT ["./server"]
EXPOSE 8080
