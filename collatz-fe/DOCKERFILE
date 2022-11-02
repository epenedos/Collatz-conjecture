
FROM golang AS collatz-fe
WORKDIR /opt/app-root/src
RUN git clone https://github.com/epenedos/Collatz-conjecture.git 
WORKDIR /opt/app-root/src/Collatz-conjecture/collatz-fe
RUN pwd
RUN ls -lisa
RUN go build
CMD  /opt/app-root/src/Collatz-conjecture/collatz-fe/collatz-fe
EXPOSE 8080
