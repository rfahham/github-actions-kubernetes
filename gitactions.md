# GITHUB ACTIONS

GITHUB ACTIONS é uma plataforma de integração contínua (CI) e entrega contínua (CD) fornecida pelo GITHUB, que permite automação, customização e execução de workflows de desenvolvimento de software diretamente no repositório GITHUB.

Neste projeto, o CI e o CD está no mesmo workflow

Criando a Service Account

```bash
kubectl -n kube-system create sa github

serviceaccount/github created
```

Criar o token

```bash
kubectl -n kube-system apply -f <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: github-secret
  annotations:
    kubernetes.io/service-account.name: github
type: kubernetes.io/service-account-token
EOF


secret/github-secret created
```

Criando uma Role Cluster Admin

```bash
kubectl create clusterrolebinding github-access --clusterrole cluster-admin --serviceaccount=kube-system:github

clusterrolebinding.rbac.authorization.k8s.io/github-access created
```

Pegar o token e salvar no KUBERNETES_TOKEN do github

```bash
kubectl -n kube-system get secrets github-secret -o json | jq -r .data.token | base64 -d
```

[Variáveis](variaveis.md)

Verificar o processo de CI/CD executando no k8s

```bash
watch -d "kubectl get pod"

Próximo passo...[Separar CI e CD](gitactions2.md)