FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/community-operator \
    USER_UID=1001 \
    USER_NAME=community-operator \
    COMMUNITY_OPERATOR_DIR=/etc/community-operator

RUN mkdir -p ${COMMUNITY_OPERATOR_DIR}/templates
COPY templates ${COMMUNITY_OPERATOR_DIR}/templates

RUN chown ${USER_UID} ${COMMUNITY_OPERATOR_DIR}
RUN chmod u=rwx,g=r ${COMMUNITY_OPERATOR_DIR}

# install operator binary
COPY build/_output/bin/community-operator ${OPERATOR}
COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
