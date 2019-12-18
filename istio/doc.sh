minikube start --docker-env HTTP_PROXY=http://127.0.0.1:1087 --docker-env HTTPS_PROXY=http://127.0.0.1:1087 --docker-env NO_PROXY=192.168.64.14 -p test # 创建test集群
kubectl config get-contexts # 查看已有集群
kubectl config use-context docker-for-desktop # 切换集群