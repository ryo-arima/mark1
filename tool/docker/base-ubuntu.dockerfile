FROM ryoarima/ubuntu-base:latest

WORKDIR /root

RUN wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1_0.0.1_amd64.deb && \
    dpkg -i mark1_0.0.1_amd64.deb

CMD [ "echo", "This is mark1" ]