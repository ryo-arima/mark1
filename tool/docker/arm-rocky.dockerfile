FROM ryoarima/rocky-arm:latest

WORKDIR /root

RUN wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-aarch64.aarch64.rpm && \
    rpm -ivh mark1-0.0.1-aarch64.aarch64.rpm

CMD [ "echo", "This is mark1" ]