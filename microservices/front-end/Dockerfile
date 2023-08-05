# FROM golang:1.20-alpine AS builder

# RUN mkdir /app

# COPY . /app

# RUN cd /app \
#     && go build -o front-end ./cmd

# FROM alpine

# RUN mkdir /app

# COPY --from=builder /app/front-end /app

# CMD /app/front-end


FROM alpine

RUN mkdir /app

WORKDIR /app

COPY front-end .
COPY cmd/web/templates/ ./cmd/web/templates/

RUN chmod +x front-end

CMD ./front-end
