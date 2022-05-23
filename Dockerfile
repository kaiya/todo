FROM alpine

WORKDIR /app
COPY ./todo /app/todo
EXPOSE 8080
ENTRYPOINT ["sh", "-c", "./todo"]
