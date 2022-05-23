FROM alpine

WORKDIR /app
COPY ./todo /app/todo
EXPOSE 8080
ARG SQLPASS

ENTRYPOINT ["sh", "-c", "./todo -sqlPass $SQLPASS"]
