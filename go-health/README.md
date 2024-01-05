### Applying `pod.yaml`
`kubectl get pods`
```
No resources found in default namespace.
```

`kubectl apply -f deployment/pod.yaml`
```
pod/go-health created
```

`kubectl get pods`
```
NAME        READY   STATUS    RESTARTS   AGE
go-health   1/1     Running   0          22s
```

`kubectl delete pod go-health`
```
pod "go-health" deleted
```


### Applying `deploy.yaml`

`kubectl get deployment`
```
No resources found in default namespace.
```

`kubectl apply -f deployment/deploy.yaml`
```
deployment.apps/go-health-container created
```

`kubectl get deployment `
```
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
go-health-container   1/1     1            1           8s
```

`kubectl get pods`
```
NAME                                   READY   STATUS    RESTARTS   AGE
go-health-container-74b684b846-ppwkf   1/1     Running   0          42s
```

`kubectl delete deployment go-health-container`
```
deployment.apps "go-health-container" deleted
```

### Applying `deploy-service.yaml`

`kubectl get services`
```
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   12h
```

Note: Our application is not running as a service.

`kubectl apply -f deployment/deploy-service.yaml `
```
deployment.apps/go-health-container created
service/go-health-service created
```

`kubectl get services `
```
NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
go-health-service   NodePort    10.97.73.231   <none>        85:30842/TCP   26s
kubernetes          ClusterIP   10.96.0.1      <none>        443/TCP        12h
```

`kubectl get deployment`
```
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
go-health-container   1/1     1            1           44s
```

`kubectl get pods `
```
NAME                                   READY   STATUS    RESTARTS   AGE
go-health-container-74b684b846-b949p   1/1     Running   0          63s
```