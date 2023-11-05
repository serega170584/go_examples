FROM ubuntu:22.04
RUN apt-get update
RUN apt-get install strace
RUN apt-get install golang-go -y
RUN apt-get -y install sudo
RUN groupadd -g 1000 sergey
RUN useradd -l -u 1000 -g sergey sergey -d /home/sergey
RUN adduser sergey sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER sergey
WORKDIR /app
#ADD . /app
#RUN sudo go mod init docker/ubuntu
#RUN sudo go build
