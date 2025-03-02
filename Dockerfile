FROM alpine:3.15

# Set environment variables
ENV PLURALITH_CI=true

ARG PLURALITH_VERSION=0.2.2
ARG PLURALITH_GRAPHING_VERSION=0.2.1
ARG TF_VERSION=1.11.0


# Install dependencies
RUN apk upgrade --no-cache && apk --no-cache add bash jq curl wget gcompat libgcc libstdc++ npm xdg-utils

# Set shell to bash to use bash array
SHELL ["/bin/bash", "-ec"]

# Install Compost for pull request commenting
RUN npm install -g @infracost/compost

# Install pluralith & terraform
RUN mkdir -p /root/Pluralith/bin/ \
  && wget https://github.com/Pluralith/pluralith-cli/releases/download/v${PLURALITH_VERSION}/pluralith_cli_linux_amd64_v${PLURALITH_VERSION} -O /usr/local/bin/pluralith \
  && wget https://github.com/Pluralith/pluralith-cli-graphing-release/releases/download/v${PLURALITH_GRAPHING_VERSION}/pluralith_cli_graphing_linux_amd64_${PLURALITH_GRAPHING_VERSION} -O /root/Pluralith/bin/pluralith-cli-graphing \
  && chmod +x /usr/local/bin/pluralith /root/Pluralith/bin/pluralith-cli-graphing \
  && wget https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip -O terraform_${TF_VERSION}_linux_amd64.zip \
  && unzip terraform_${TF_VERSION}_linux_amd64.zip -d /usr/local/bin \
  && rm terraform_${TF_VERSION}_linux_amd64.zip \
  && curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

ENTRYPOINT [ "pluralith" ]

