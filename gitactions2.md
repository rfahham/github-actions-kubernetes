# Separar CI do CD no GITACTIONS

Se quisermos escalar a quantidade de réplicas o deploy não seja realizado e vice-versa

Criar 2 repositórios diferentes

Um para a Infra (CI)

- infra

Outra para o deploy (CD)

- code

Alterar as variáveis

> Code

Secrets

DOCKER_PASSWORD - 

INFRA_TOKEN - token .ssh

Variables

DOCKER_REPO - rfahham/goapp:v1

DOCKER_USERNAME - rfahham

INFRA_USERNAME - rfahham


> Infra

Secrets

- KUBERNETES_TOKEN - 

variables

- K8S_API_SERVER - https://3.81.100.103:6443 - O IP do CLUSTER