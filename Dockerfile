
FROM golang AS collatz-fe
RUN pwd
RUN chmod +rwx /src
RUN ls -lisa
WORKDIR /src
RUN git clone https://github.com/epenedos/Collatz-conjecture.git 
WORKDIR /src/Collatz-conjecture/collatz-fe
RUN pwd
RUN ls -lisa
RUN go build
CMD  /src/Collatz-conjecture/collatz-fe/collatz-fe
EXPOSE 8080
