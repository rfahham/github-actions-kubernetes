# Buidando a aplicação MANUALMENTE

```bash
cd app
docker build -t rfahham/goapp:v1 .
```

Verificando a imagem buidada

```bash
docker images
```

Logando no Docker Hub

```bash
docker login
```

Upload para o registry (dockerHUB)

```bash
docker push rfahham/goapp:v1
```

Confirmar no docker HUB

No Cluster Kubernetes


```bash
kubectl get nodes
```

Criar o namespace

```bash
kubectl create ns goapp
```

Mudar para o namespace goapp

```bash
kubens goapp
```

Criar o Deployment

```bash
kubectl create deployment goapp --image=rfahham/goapp:v1
```

Verificar o Deployment

```bash
kubectl get pods
```

Expor/create o Deployment

```bash
kubectl expose deployment goapp port=8080

ou 

kubectl create service clusterip goapp --tcp=8080:8080
```

Verificar se o serviço está up

```bash
kubectl get svc
```

Criar o Ingress (já está cadastrado no DNS)

```bash
kubectl create ingress goapp --rule="<dominio>/*=goapp=8080"
```

Para fazer alteração na aplicação, editar o arquivo, salvar, buidar e fazer o push com a nova versão

```bash
vi app/server.go
docker build -t rfahham/goapp:v2 .
docker push rfahham/goapp:v2
```

Atualizar a aplicação em kubernetes

```bash
kubectl set image deployment/goapp rfahham/goapp:v2
```

Verificar as imagens

```bash
kubectl get pod -o custom-coluns=CONTAINER:.spec.container[0].name, IMAGE:.spec.container[0].image
```

```bash
kubectl get pod
```

Editar o deployment

```bash
kubectl edit deployments.goapp
```

Visualizando o manifesto deployments

```bash
kubectl create deployment goapp --image=rfahham/goapp:v1 --dry-run=client -o yaml
```

Aplicando 

```bash
kubectl create deployment goapp --image=rfahham/goapp:v1 --dry-run=client -o yaml | kubectl apply -f -
```