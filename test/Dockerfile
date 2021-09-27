FROM ubuntu

# WORKDIR /dev
RUN apt update
RUN apt install curl -y
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs  
RUN node -v 
RUN npm -v
RUN npm i -g barego
COPY ./baretest.sh .
RUN bare
RUN chmod +x ./baretest.sh
RUN ./baretest.sh