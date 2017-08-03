
    # Pointing Docker to minikube daemon
    eval $(minikube docker-env)

    # Web
    GOOS=linux go build -o target/hello hello.go
    docker build -t jbugman/hello-kube-web:v1 -f web/Dockerfile web
    kubectl run hello-web --image=jbugman/hello-kube-web:v1
    kubectl expose deployment hello-web --name=web --port=8080

    # Deployment rollout
    kubectl set image deployment/hello-web hello-web=jbugman/hello-kube-web:v2

    # Frontend
    docker build -t jbugman/hello-kube-front:latest -f frontend/Dockerfile frontend
    kubectl expose deployment hello-web --type=NodePort --name=frontend --port=80
