FROM debian:jessie

COPY bin/practice-server .

ARG LISTENPORT=8080

ARG CACHEADDRESS=localhost
ARG CACHEPORT=8081

ENV LISTENPORT=${LISTENPORT}
ENV CACHEADDRESS=${CACHEADDRESS}
ENV CACHEPORT=${CACHEPORT}

CMD ["/bin/bash", "-c", "/practice-server -p=${LISTENPORT} --cache-endpoint=${CACHEADDRESS} --cache-port=${CACHEPORT}"]

EXPOSE ${LISTENPORT}