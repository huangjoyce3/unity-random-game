FROM alpine
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY gateway /gateway
EXPOSE 443
ENTRYPOINT ["/gateway"]