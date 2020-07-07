#!/bin/bash
export NGROK_DOMAIN=ngrok.kinjuhui.com

mkdir -p ./certs

openssl genrsa -out ./certs/rootCA.key 2048
openssl req -x509 -new -nodes -key ./certs/rootCA.key -subj "/CN=$NGROK_DOMAIN" -days 5000 -out ./certs/rootCA.pem
openssl genrsa -out ./certs/server.key 2048
openssl req -new -key ./certs/server.key -subj "/CN=$NGROK_DOMAIN" -out ./certs/server.csr
openssl x509 -req -in ./certs/server.csr -CA ./certs/rootCA.pem -CAkey ./certs/rootCA.key -CAcreateserial -out ./certs/server.crt -days 5000

cp ./certs/rootCA.pem ./assets/client/tls/ngrokroot.crt
cp ./certs/server.crt ./assets/server/tls/snakeoil.crt
cp ./certs/server.key ./assets/server/tls/snakeoil.key