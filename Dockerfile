FROM ubuntu 
RUN  apt update && apt upgrade -y && apt install golang-go -y
RUN apt update && \
    apt install -y sudo && \
    addgroup --gid $GID sanket && \
    adduser --uid $UID --gid $GID --disabled-password --gecos "" nonroot && \
    echo 'nonroot ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

USER sanket
WORKDIR /home/sanket
COPY ./main.go .
COPY ./main1.go .
ENTRYPOINT [ "go" ,"run" ] 
CMD [ "main.go" ]


