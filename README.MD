# Unosquare Challenge

## Installation
* Clone repo
```
git clone git@github.com:PerezBautistaAntonioDeJesus/unosquare-challenge.git
```


## Requirements

* Docker
* Minikube
* Kubectl


## Prepare the environment

* To avoid publishing our images to a registry, we are going to build the images to be reached from minikube.

* run:
```
 minikube docker-env
```

* this will show you the environment variables we need to export.
* run:
``` 
eval $(minikube -p minikube docker-env)
```

## Initialize minikube tunnel

* In order to access the services of your cluster locally with minikube, we must initialize the tunnel to expose the ports from our host with the following command.

````
minikube tunnel
````


## build images

### building the front end image.
* access the frontend directory.

```
cd frontend
````

 * build the image wiht the name `unosquare-challenge-web`
```
docker build -t unosquare-challenge-web .
````

* go back to the root directory
```
cd ..
````

### building the api service image.
* access the unosquare-challenge-service directory

```
cd unosquare-challenge-service
```

* build the image wiht the name `unosquare-challenge`
```
docker build -t unosquare-challenge .
````

* apply the kubernetes manifest with the following command.

```
kubectl apply -f deployment.yaml
```

* then run
````
kubectl get all
````

* you should see something similar to this.
```
NAME                            READY   STATUS    RESTARTS   AGE
pod/unosquare-9fc8d7f68-9r5t7   1/1     Running   0          81s
pod/unosquare-9fc8d7f68-gv6p9   1/1     Running   0          81s
pod/unosquare-9fc8d7f68-svd8v   1/1     Running   0          81s
pod/unosquare-web               1/1     Running   0          81s

NAME                    TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/kubernetes      ClusterIP      10.96.0.1      <none>        443/TCP          46h
service/unosquare       LoadBalancer   10.97.17.250   127.0.0.1     8080:31078/TCP   81s
service/unosquare-web   LoadBalancer   10.106.76.19   127.0.0.1     8082:31360/TCP   81s

NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/unosquare   3/3     3            3           81s

NAME                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/unosquare-9fc8d7f68   3         3         3       81s
```

-----

## Web application :shipit:
http://127.0.0.1:8082

## API Documentation

http://127.0.0.1:8080/docs