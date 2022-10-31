
FROM golang AS collatz-fe
WORKDIR /src
RUN git clone https://github.com/epenedos/Collatz-conjecture.git /src
WORKDIR /src/Collatz-conjecture/collatz-fe
RUN pwd
RUN ls -lisa
RUN go build
CMD  /src/Collatz-conjecture/collatz-fe/collatz-fe
EXPOSE 8080
