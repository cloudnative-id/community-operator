# community-operator
The Community Operator provides Kubernetes native deployment and management of your community. The purpose of this project is to simplify and automate community management in top of Kubernetes clusters.

### Design Decision
WIP

### Feature
- [x] weekly controller
- [ ] information controller
- [ ] meetup controller
- [x] telegram dispatcher
- [ ] twitter dispatcher
- [ ] facebook dispatcher

### Developing community-operator
This operator build based on [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/), to build this operator, you need [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/).

#### Running community-operator
- export variable need for community-operator
```
export TELEGRAM_TOKEN=xxx
export TELEGRAM_CHATID=yyy
```
- deploy CRDs
```
kubectl apply -f deploy/crds/community.io_weeklies_crd.yaml
```
- run community-operator locally
```
operator-sdk run local
```
- deploy some example weekly
```
kubectl apply -f examples/
```

