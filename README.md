
    # Pointing Docker to minikube daemon
    eval $(minikube docker-env)

    # Web
    GOOS=linux go build -o web/target/hello web/hello.go
    docker build -t jbugman/hello-kube-web:latest -f web/Dockerfile web
    kubectl run hello-web --image=jbugman/hello-kube-web:v1
    kubectl expose deployment hello-web --name=web --port=8080

    # Backend
    GOOS=linux go build -o backend/target/hello backend/backend.go
    docker build -t jbugman/hello-kube-backend:latest -f backend/Dockerfile backend
    kubectl run hello-backend --image=jbugman/hello-kube-backend:v1
    kubectl expose deployment hello-backend --name=backend --port=80 --target-port=8080

    # Deployment rollout
    kubectl set image deployment/hello-web hello-web=jbugman/hello-kube-web:v2

    # Frontend
    docker build -t jbugman/hello-kube-front:latest -f frontend/Dockerfile frontend
    kubectl run hello-front --image=jbugman/hello-kube-front:v1
    kubectl expose deployment hello-front --type=NodePort --name=frontend --port=80
