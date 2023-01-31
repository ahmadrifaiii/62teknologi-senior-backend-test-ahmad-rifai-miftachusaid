###########################################################################
# Stage 1 Start
###########################################################################
FROM golang AS build-golang


RUN export GO111MODULE=on
RUN export GOPROXY=direct
RUN export GOSUMDB=off

################################
# Build Service:
################################
WORKDIR /usr/share/62tech/62service

COPY  . .

RUN make deploy

###########################################################################
# Stage 2 Start
###########################################################################
FROM ubuntu:20.04

# Change Repository ke kambing.ui:
RUN sed -i 's*archive.ubuntu.com*kambing.ui.ac.id*g' /etc/apt/sources.list

RUN apt-get update

RUN apt-get install -y ca-certificates

# Copy Binary
COPY --from=build-golang /usr/share/62tech/62service/bin /usr/share/62tech/62service/bin/

WORKDIR /usr/share/62tech/62service

# Create group and user to the group
RUN groupadd -r 62service && useradd -r -s /bin/false -g 62service 62service

# Set ownership golang directory
RUN chown -R 62service:62service /usr/share/62tech/62service

# Make docker container rootless
USER 62service

# EXPOSE 8080

# ENTRYPOINT [ "./service" ]
# CMD [ "./rest" ]