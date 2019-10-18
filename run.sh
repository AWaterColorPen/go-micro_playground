docker build -t gomicrok8s:1.0.8 -f k8s/Dockerfile .
kubectl apply -f k8s/gomicrok8s.yaml