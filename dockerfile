FROM scratch
LABEL maintainer "ladislav.lisy@seznam.cz"

COPY payroll-maestro-server /home/payroll-maestro-server
COPY templates /home/templates
