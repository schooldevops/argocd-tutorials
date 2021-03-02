# ArgoCD 디플로이 

이제는 ArgoCD 를 이용하여 kubernetes 에 배포를 수행할 것이다. 

그러기 위해서는 기본적으로 2가지 작업을 수행해 주어야한다. 

- Deployment: Kubernetes 에 어플리케이션을 배포하기 위해서는 Deployment 를 통해서 배포를 수행하게 된다. 버전관리 등 다양한 이점이 있다. 
- Service: Service 는 외부 접속을 위한 연결 정의를 수행하는 매니페스트이다. 

## Deployment 작성하기. 

greetweb-deploy.yml 파일을 다음과 같이 작성한다. 

```yml
apiVersion: apps/v1
kind: Deployment 
metadata:
  labels:
    app: greet
  name: greet 
spec:
  replicas: 2
  selector:
    matchLabels:
      app: greet
  template:
    metadata:
      labels:
        app: greet 
    spec:
      containers:
        - image:  unclebae/gogreet:v1.0
          name:  unclebae/gogreet
```

## Service 작성하기 

서비스를 생성한다. 

```yml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: greet
  name: greet
  namespace: default 
spec:
  ports:
  - port: 9999
    protocol: TCP
    targetPort: 8080
  selector:
    app: greet
  type: NodePort

```

## 위 작성한 코드를 모두 github으로 올린다. 

## ArgoCD 접속하기. 

https://localhost:31602/

으로 지난번에 argocd 를 설치 했었다. 이제는 argocd 로 배포를 수행해 보자. 

