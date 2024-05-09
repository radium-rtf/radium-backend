FROM postgres:16

RUN apt update && apt upgrade -y
RUN apt install -y postgresql-server-dev-16
RUN apt install -y git
RUN apt install -y systemtap-sdt-dev
RUN apt-get install -y build-essential
RUN apt-get install -y make && apt -y install pgxnclient

RUN git clone https://github.com/postgrespro/rum && \
    cd rum &&  \
    make USE_PGXS=1 &&  \
    make USE_PGXS=1 install \
