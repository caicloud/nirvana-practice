FROM debian:jessie

ARG PORT=8081

COPY bin/practice-cache .

ENV PORT=${PORT}

CMD ["/bin/bash", "-c", "/practice-cache -p=${PORT}"]

EXPOSE ${PORT}
