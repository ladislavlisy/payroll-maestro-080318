FROM scratch
LABEL maintainer "ladislav.lisy@seznam.cz"

COPY . /
ENTRYPOINT [ "payroll-maestro-server" ]