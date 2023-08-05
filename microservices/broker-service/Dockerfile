# FROM golang:1.20-alpine AS builder

# RUN mkdir /app

# COPY . /app

# RUN cd /app \
#     && go build -o broker-service ./cmd/api

# FROM alpine

# RUN mkdir /app

# COPY --from=builder /app/broker-service /app

# CMD /app/broker-service


FROM alpine

RUN mkdir /app

WORKDIR /app

COPY broker-service .

RUN chmod +x broker-service

CMD ./broker-service
