FROM ubuntu:latest

RUN apt update
RUN apt upgrade
#RUN unminimize -y

RUN apt install zsh  git sudo man    -y
RUN apt install golang -y

RUN useradd --create-home --shell /bin/zsh vscode
RUN adduser vscode sudo 

RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >>/etc/sudoers
RUN echo "vscode:vscode" | chpasswd


COPY ./.zshrc /home/vscode