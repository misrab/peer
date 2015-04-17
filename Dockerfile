FROM zeromq/zeromq
MAINTAINER <Misrab M. Faizullah-Khan> <faizullah.misrab@gmail.com>


RUN apt-get update && apt-get -y upgrade
# install c compiler etc
RUN apt-get install -y git build-essential emacs
#RUN gcc-v && make -v
# install golang from source
RUN git clone https://go.googlesource.com/go && cd go && git checkout go1.4.1
RUN cd go/src && ./all.bash
# set go path
RUN echo 'export GOROOT=/go\nexport PATH=$PATH:$GOROOT/bin' >> ~/.bashrc && . ~/.bashrc
RUN mkdir /home/go && mkdir /home/go/src && mkdir /home/go/src/github.com && mkdir /home/go/src/github.com/misrab
RUN echo 'export GOPATH=/home/go PATH=$PATH:$GOPATH' >> ~/.bashrc && . ~/.bashrc

# get my code
RUN cd /home/go/src/github.com/misrab && git clone https://github.com/misrab/peer.git


# this is very sensitive to env vars for obscure reasons
CMD ["sh", "-c", ". ~/.bashrc &&  echo moo && echo ${HOME} && echo ${GOPATH} && cd /home/go/src/github.com/misrab/peer && git pull && go clean && go get && go test"]
