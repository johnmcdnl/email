FROM debian:stretch-slim

MAINTAINER John McDonnell

RUN apt-get update
RUN apt-get install -y postfix
RUN apt-get install -y dovecot-pop3d
RUN apt-get install -y dovecot-imapd
RUN apt-get install -y syslog-ng


# TODO users from ENV
ENV USERNAME johnmcdnl
ENV PASSWORD password
ENV HOSTNAME mailserver


# Configure postfix
RUN useradd -s /bin/bash ${USERNAME}
RUN mkdir /var/spool/mail/${USERNAME}
RUN chown ${USERNAME}:mail /var/spool/mail/${USERNAME}

RUN postconf -e myhostname=${HOSTNAME}
RUN postconf -e mydestination="mail.${HOSTNAME}, ${HOSTNAME}, localhost.localdomain, localhost"
RUN postconf -e mail_spool_directory="/var/spool/mail/"
RUN postconf -e mailbox_command=""

# Configure dovecot
RUN echo "mail_location = mbox:/var/spool/mail/%u"  >> "/etc/dovecot/dovecot.conf"
RUN echo "disable_plaintext_auth = no"              >> "/etc/dovecot/dovecot.conf"
RUN echo "auth_mechanisms = plain login"            >> "/etc/dovecot/dovecot.conf"
RUN echo "mail_privileged_group = mail"             >> "/etc/dovecot/dovecot.conf"
RUN echo "protocols = imap pop3"                    >> "/etc/dovecot/dovecot.conf"


RUN echo "${USERNAME}:{PLAIN}${PASSWORD}:$UID:$GID::$HOME" > users
RUN mv users /etc/dovecot/
RUN sed -i 's/#!include auth-passwdfile.conf.ext/!include auth-passwdfile.conf.ext/g' /etc/dovecot/conf.d/10-auth.conf


EXPOSE 25 143
CMD ["sh", "-c", "service dovecot start ; service syslog-ng start ; service postfix start ; tail -F /var/log/mail.log"]