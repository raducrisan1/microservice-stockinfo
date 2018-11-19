FROM alpine:3.8
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY ./microservice-stockinfo /app/
WORKDIR /app
EXPOSE 3001
CMD ["./microservice-stockinfo"]