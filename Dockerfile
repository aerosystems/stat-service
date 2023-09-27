FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./stat-service/stat-service.bin /app

# Run the server executable
CMD [ "/app/stat-service.bin" ]