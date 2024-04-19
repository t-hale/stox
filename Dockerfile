FROM ubuntu:latest

ARG port
ENV PORT=$port

RUN apt-get update
RUN apt-get install -y ca-certificates

WORKDIR /app

COPY server .

EXPOSE 0.0.0.0:$PORT:$PORT

CMD /app/server -http-port 8080 -host 0.0.0.0
#CMD nc -l 8080