FROM ubuntu:20.04

RUN apt update
RUN apt install -y openssh-server
RUN bash -c "echo root:toor | chpasswd"
RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config
RUN mkdir /run/sshd

RUN apt install -y build-essential make golang libpam-dev pamtester vim ca-certificates git sudo
RUN echo 'auth required pam_websso.so' >> /etc/pam.d/sshd
RUN sed -i 's/^ChallengeResponseAuthentication no$/ChallengeResponseAuthentication yes/g' /etc/ssh/sshd_config

RUN ssh-keygen -A

RUN groupadd --gid 1000 workers \
    && useradd --uid 1000 --gid workers --shell /bin/bash --create-home worker \
    && adduser worker sudo

RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

WORKDIR /home/worker/work

USER 1000

ENTRYPOINT sudo service ssh restart && bash