FROM ryoarima/rocky-base:latest

WORKDIR /root

RUN wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-x86_64.x86_64.rpm && \
    rpm -ivh mark1-0.0.1-x86_64.x86_64.rpm

CMD [ "echo", "This is mark1" ]