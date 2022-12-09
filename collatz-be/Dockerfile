FROM image-registry.openshift-image-registry.svc:5000/openshift/golang:latest AS collatz-be
WORKDIR /opt/app-root/src
RUN git clone https://github.com/epenedos/Collatz-conjecture.git 
WORKDIR /opt/app-root/src/Collatz-conjecture/collatz-be
RUN go build
CMD  /opt/app-root/src/Collatz-conjecture/collatz-be/collatz-be
EXPOSE 8081
