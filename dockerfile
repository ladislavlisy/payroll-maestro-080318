FROM scratch
LABEL maintainer "ladislav.lisy@seznam.cz"

COPY /go/src/github.com/ladislavlisy/payroll-maestro-080318/payroll-maestro-server /
COPY /go/src/github.com/ladislavlisy/payroll-maestro-080318/server/templates /
ENTRYPOINT ["/payroll-maestro-server"]

RUN chmod +x /payroll-maestro-server
