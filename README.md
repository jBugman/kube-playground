
    # Pointing Docker to minikube daemon
    eval $(minikube docker-env)

    # Web
    GOOS=linux go build -o web/target/hello web/hello.go
    docker build -t jbugman/hello-kube-web:latest -f web/Dockerfile web

    # Backend
    GOOS=linux go build -o backend/target/hello backend/backend.go
    docker build -t jbugman/hello-kube-backend:latest -f backend/Dockerfile backend

    # Frontend
    docker build -t jbugman/hello-kube-front:latest -f frontend/Dockerfile frontend

    # Declarative deploy
    kubectl apply -f config/

    # Manual deployment rollout
    kubectl set image deployment/hello-web hello-web=jbugman/hello-kube-web:v2