x509 error when using HTTPS inside a Docker container
Khash Sajadi 19 MARCH 2015
If you are getting a 'x509: failed to load system roots and no roots provided' error when you run an HTTPS from inside a container, this might help.

This issue can happen if the base image used to build your Docker container does not have the CA certificates. Here is how it can be reproduced:

FROM ubuntu:14.04.1
CMD curl https://www.google.com
Running this container might generate something like this:

x509: failed to load system roots and no roots provided
The reason for this is that the base image (ubuntu:14.04.1 in this case) does not have the root CA certificates installed. The solution is easy enough. Install them as part of your image build:

FROM ubuntu:14.04.1

RUN apt-get update
RUN apt-get install -y ca-certificates

CMD curl https://www.google.com

We have a small system running in AWS as a CentOS 7 image. It has a few containers that we’re using to host a few Golang API proxies. We migrated a customers API proxy that was running on the local VM into a container, and spun it up. Upon testing, we ran into the following error:

1
x509: failed to load system roots and no roots provided
We get that failure when trying to connect to an HTTPS endpoint (remote API that we’re proxying to Asterisk).

Figured it had to do with the fact we were using a scratch disk to build the container image, and that there were no certs loaded. Did some Googling and found some people with similar problems, but their solutions didn’t work for us on our CentOS 7 host system.

Then I thought maybe there was some issue with following a symlink as the source since we were loading in the ca-bundle.crt file as a volume. I didn’t test enough to determine if that was the issue (it probably wasn’t), but this post gave me a hint:

https://github.com/docker/docker/issues/5157#issuecomment-69325677

So we did the following:

1
docker run -d -p 8085:8085 -v /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem:/etc/ssl/certs/ca-certificates.crt [etc...]
After linking that file and mounting it in the container, all was well. I suspect it’s the path to the ca-certificates.crt that was the real trick.

I created a docker container for a golang app, based on debian 7 image. Everything works well until it visits https. The error is x509: failed to load system roots and no roots provided.

I thought there might be something that should be configured for golang. But actually, it is nothing about golang. It's about SSL certificates. For debian, /etc/ssl/certs/ca-certificates.crt is the root certificates. But the offical debian 7 image does not contain it. Installing ca-certificates via apt-get solved the problem. (You can manually copy /etc/ssl/certs/ca-certificates.crt from another linux OS too.)

Dockerfile:

FROM debian:7
RUN apt-get update \
 && apt-get install -y --force-yes --no-install-recommends \
      apt-transport-https \
      curl \
      ca-certificates \
 && apt-get clean \
 && apt-get autoremove \
 && rm -rf /var/lib/apt/lists/*

// Possible certificate files; stop after finding one.
 var certFiles = []string{
+	"/etc/ssl/certs/ca-certificates.crt",                // Debian/Ubuntu/Gentoo etc.
+	"/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem", // CentOS/RHEL 7
+	"/etc/pki/tls/certs/ca-bundle.crt",                  // Fedora/RHEL 6
+	"/etc/ssl/ca-bundle.pem",                            // OpenSUSE
+	"/etc/pki/tls/cacert.pem",                           // OpenELEC
 }
