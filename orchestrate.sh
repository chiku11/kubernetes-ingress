kubectl delete all --all -n demo
kubectl delete ingress/demo-ingress -n demo
cd person
docker build -t personserver:1.0 .  
kubectl apply -f person.deploy.yml
kubectl apply -f person.svc.yml
cd ..
cd location
docker build -t locationserver:1.0 .  
kubectl apply -f location.deploy.yml
kubectl apply -f location.svc.yml
cd ..
kubectl apply -f ingress-controller.yml
kubectl apply -f ingress-config.yml