FROM alpine:3.18

WORKDIR /app

ADD https://github.com/Qolorerr/OperatorServiceAPI/releases/download/v1.0.0/operator_text_channel /app/operator_text_channel

# ИЛИ если бинарник уже в локальной папке:
# COPY operator_text_channel /app/

RUN chmod +x /app/operator_text_channel

RUN apk add --no-cache tzdata

ENTRYPOINT ["/app/operator_text_channel"]
