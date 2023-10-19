# Base image, for binary compilation
FROM golang:1.20 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/apicrud

# Now copy it into distroless base image
FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=build /go/bin/apicrud /app/apicrud
CMD ["/app/apicrud"]
