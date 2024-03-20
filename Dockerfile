FROM golang:alpine AS build-stage

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/bin/portolio-app

# FROM build-stage AS test-stage

# RUN make test


FROM scratch AS release-stage

WORKDIR /

COPY --from=build-stage /tmp/bin/portolio-app /portolio-app
COPY --from=build-stage /app/assets /assets

EXPOSE 3000

# USER nonroot:nonroot

ENTRYPOINT [ "/portolio-app" ]