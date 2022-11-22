# Cluster Kubernetes
```
https://kind.sigs.k8s.io/
```

## enviar para o cluster kubernetes
```
kubectl apply -f k8s
```
## cluster kubernetes status deploy
```
kubectl get deploy
```

## cluster kubernetes services
```
kubectl get svc
```

## cluster kubernetes context
```
kubectl cluster-info --context kind-gitops
```

## cluster kubernetes get all
```
kubectl get all -n <name>
```

# Kustomize [files deployment]
```
https://kustomize.io/
```

# Agente mudanças repositórios [ArgoCD]
```
https://argoproj.github.io/cd/
```

# Get password [ArgoCD] - dFEdIRthZlX4SNCk
```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

# Run Port forward [ArgoCD]
```
kubectl port-forward svc/argocd-server -n argocd 8080:443
```
