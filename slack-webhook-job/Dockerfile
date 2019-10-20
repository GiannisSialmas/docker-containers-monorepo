FROM golang:1.12-alpine AS build

# Metadata
ENV GOOS linux
ENV GOARCH amd64
# Turn off CGO since that can result in dynamic links to libc/libmusl.
ENV CGO_ENABLED 0 

WORKDIR /app 

#Copy and build the app
COPY main.go main.go
RUN go build -ldflags="-w -s" -o ./slack-notify ./main.go

#The certificates are needed cause scratch will keep giving "x509: certificate signed by unknown authority" on requests
RUN apk --no-cache add ca-certificates  


FROM scratch
COPY --from=build /app/slack-notify /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/slack-notify"]  