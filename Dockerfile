FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs
RUN mkdir /app/certs

COPY ./stat-service.bin /app

# Run the server executable
CMD [ "/app/stat-service.bin" ]