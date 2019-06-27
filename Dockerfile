FROM golang as build

WORKDIR /build
COPY ./ /build
RUN make build-linux64

FROM scratch

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /build/bin/redirector /

ENV PORT=8080
EXPOSE 8080

CMD ["/redirector"]
