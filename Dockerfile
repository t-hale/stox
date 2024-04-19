FROM ubuntu:latest

ARG port
ENV PORT=$port

RUN apt update
RUN apt install -y apt-transport-https ca-certificates curl gnupg
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN echo "deb https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN apt update
RUN apt install -y google-cloud-sdk

WORKDIR /app

COPY server .

EXPOSE 0.0.0.0:$PORT:$PORT

CMD /app/server -http-port 8080 -host 0.0.0.0
#CMD nc -l 8080