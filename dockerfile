FROM scratch
LABEL maintainer "ladislav.lisy@seznam.cz"

COPY . /

CMD [ "payroll-maestro-server" ]