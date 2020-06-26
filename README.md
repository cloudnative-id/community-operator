# community-operator
The Community Operator provides Kubernetes native deployment and management of your community. The purpose of this project is to simplify and automate community management in top of Kubernetes clusters.

### Feature
#### API
- [x] weekly API & controller
- [x] announcement API & controller
- [x] meetup API & controller

#### Dispatcher
- [x] telegram dispatcher
- [ ] support for multiple telegram dispatcher
- [ ] twitter dispatcher
- [ ] support for multiple twitter dispatcher
- [ ] facebook dispatcher
- [ ] support for multiple facebook dispatcher

#### Installation
- [x] helm chart

### Developing community-operator
This operator build based on [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/). To build this operator, you need [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/).

#### Running community-operator
- export variable need for community-operator
```
export TELEGRAM_TOKEN=xxx
export TELEGRAM_CHATID=yyy
```
- deploy CRDs
```
kubectl apply -f deploy/crds/community.io_weeklies_crd.yaml
kubectl apply -f deploy/crds/community.io_meetups_crd.yaml
```
- run community-operator locally
```
operator-sdk run local
```
- deploy some example
```
kubectl apply -f examples/weekly/example.yaml
kubectl apply -f examples/meetup/example.yaml
```

#### Building community-operator
To build community-operator image, you can use community-operator from root of the project
```
operator-sdk build cloudnativeid/community-operator:0.0.1
```

#### Installing community-operator via helm
Please read README.md in charts folder for more information.
```
helm install ./charts --name-template community-operator --set-string telegram.chatid="-1234556" --set-string telegram.token="12354:asdaADASFD"
```

to upgrade
```
helm upgrade community-operator ./charts --set-string telegram.chatid="-1234556" --set-string telegram.token="12354:asdaADASFD"
```
