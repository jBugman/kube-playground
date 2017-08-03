
    GOOS=linux go build -o target/hello hello.go
    eval $(minikube docker-env)
    docker build -t jbugman/hello-kube-web:v1 .

    kubectl run hello-web --image=jbugman/hello-kube-web:v1
    kubectl expose deployment hello-web --type=NodePort --name=web --port=8080
    minikube service web

    kubectl set image deployment/hello-web hello-web=jbugman/hello-kube-web:v2
