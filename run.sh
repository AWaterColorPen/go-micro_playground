docker build -t gomicrok8s:1.2.0 -f k8s/Dockerfile .
kubectl apply -f k8s/gomicrok8s.yaml