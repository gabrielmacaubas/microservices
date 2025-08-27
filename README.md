# MSGo - Execução dos Microsserviços

## Requisitos

- Docker
- Minikube (Kubernetes local)
- kubectl

---

## Executando com Docker Compose

1. **Construa as imagens dos microsserviços:**
   ```sh
   docker build -t order:latest ./microservices/order
   docker build -t payment:latest ./microservices/payment
   docker build -t shipping:latest ./microservices/shipping
   ```

2. **Suba os containers (caso possua um arquivo `docker-compose.yml`):**
   ```sh
   docker-compose up --build
   ```

3. **Para parar os containers:**
   ```sh
   docker-compose down
   ```

---

## Executando com Kubernetes (Minikube)

1. **Inicie o Minikube:**
   ```sh
   minikube start
   ```

2. **Configure o Docker para usar o registry do Minikube:**
   ```sh
   minikube docker-env | Invoke-Expression
   ```

3. **Construa as imagens dos microsserviços:**
   ```sh
   docker build -t order:latest ./microservices/order
   docker build -t payment:latest ./microservices/payment
   docker build -t shipping:latest ./microservices/shipping
   ```

4. **Implante os serviços e banco de dados:**
   ```sh
   kubectl apply -f microservices/k8s/postgres/
   kubectl apply -f microservices/k8s/order/
   kubectl apply -f microservices/k8s/payment/
   kubectl apply -f microservices/k8s/shipping/
   ```

5. **Verifique os pods e serviços:**
   ```sh
   kubectl get pods
   kubectl get services
   ```

6. **Para acessar um serviço via Minikube:**
   ```sh
   minikube service order-service --url
   minikube service payment-service --url
   minikube service shipping-service --url
   ```

7. **Para parar o Minikube:**
   ```sh
   minikube stop
   ```

---

## Informações de Banco de Dados

- **Host:** postgres-service
- **Porta:** 5432
- **Usuário:** msgo
- **Senha:** msgo123
- **Banco:** msgo_db

---

## Portas dos Serviços

| Serviço   | Porta  |
|-----------|--------|
| Order     | 50051  |
| Payment   | 50052