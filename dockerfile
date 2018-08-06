FROM scratch
LABEL maintainer "ladislav.lisy@seznam.cz"

COPY . /
RUN chmod +x payroll-maestro-server

CMD [ "payroll-maestro-server" ]