FROM alpine

RUN mkdir /app

WORKDIR /app

COPY authentication-service .

RUN chmod +x authentication-service

CMD ./authentication-service
