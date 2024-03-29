FROM scratch

WORKDIR /app

ADD ./main /app
ADD ./certs /app/certs

EXPOSE 3000

ENTRYPOINT [ "./main" ]