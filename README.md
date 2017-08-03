
    eval $(minikube docker-env)

    GOOS=linux go build -o target/hello hello.go
    docker build -t jbugman/hello-kube-web:v1 -f web/Dockerfile web

    kubectl run hello-web --image=jbugman/hello-kube-web:v1
    kubectl expose deployment hello-web --type=NodePort --name=web --port=8080
    minikube service web

    kubectl set image deployment/hello-web hello-web=jbugman/hello-kube-web:v2

    docker build -t jbugman/hello-kube-front:v1 -f frontend/Dockerfile frontend
