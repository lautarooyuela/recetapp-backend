FROM scratch

# Se setea variable de entorno
ARG RECENV=P
ENV RECENV $RECENV

# Se instala app
ADD main /
ADD .env /
COPY Resources/ /Resources

# Se copia certificado de maquina host
#COPY ca-certificates.crt /etc/ssl/certs/

# Se habilita puerto
EXPOSE 4000

# Se ejecuta solo cuando se corre "docker run"
ENTRYPOINT ["/main"]