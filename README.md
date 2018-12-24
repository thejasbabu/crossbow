# Crossbow

This project is implemetation of FullStack CRD which deploys an entire application in kubernetes

## Steps

`minikube start` to start the minikube instance  
`make install` installs the crd to the minikube cluster  
`make run` runs the manager  

To apply the definition  
`kubectl apply definition/fullstack.yaml`