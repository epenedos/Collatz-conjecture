
FROM golang AS collatz-fe
WORKDIR /src
RUN ls -lisa
RUN git clone https://github.com/epenedos/Collatz-conjecture.git 
WORKDIR /src/Collatz-conjecture/collatz-fe
RUN pwd
RUN ls -lisa
RUN go build
CMD  /src/Collatz-conjecture/collatz-fe/collatz-fe
EXPOSE 8080