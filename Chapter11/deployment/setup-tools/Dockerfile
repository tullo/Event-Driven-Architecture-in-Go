FROM alpine as terraform-installer

ARG TERRAFORM_VERSION="1.4.2"

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip; \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip; \
    mv terraform /usr/local/bin

FROM alpine as helm-installer

ARG HELM_VERSION="3.11.2"

RUN wget https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz; \
    tar -zxvf helm-v${HELM_VERSION}-linux-amd64.tar.gz; \
    mv linux-amd64/helm /usr/local/bin

FROM alpine as k8s-installer

RUN apk update; \
    apk add k9s

RUN apk add curl; \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"; \
    chmod +x kubectl; \
    mv kubectl /usr/local/bin

FROM alpine

COPY --from=terraform-installer /usr/local/bin/terraform /usr/local/bin
COPY --from=helm-installer /usr/local/bin/helm /usr/local/bin
COPY --from=k8s-installer /usr/bin/k9s /usr/local/bin
COPY --from=k8s-installer /usr/local/bin/kubectl /usr/local/bin

RUN apk update; \
    apk add postgresql-client make git docker openrc

CMD "sh"
