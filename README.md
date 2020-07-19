# community-operator
The Community Operator provides Kubernetes native deployment and management of your community. The purpose of this project is to simplify and automate community management in top of Kubernetes clusters.

Community-operator currently watch CRDs on the same namespace as the community-operator deployed.

### Feature
#### API
- [x] weekly API & controller
- [x] announcement API & controller
- [x] meetup API & controller

#### Dispatcher
- [x] telegram dispatcher
- [ ] support for multiple telegram dispatcher
- [x] twitter dispatcher
- [ ] support for multiple twitter dispatcher
- [ ] facebook dispatcher
- [ ] support for multiple facebook dispatcher

#### Installation
- [x] helm chart

#### Other
- [ ] support multi namespace CRDs

### Developing community-operator
This operator build based on [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/). To build this operator, you need [operator-sdk](https://sdk.operatorframework.io/docs/install-operator-sdk/).

#### Running community-operator
- export variable need for community-operator
```
export TELEGRAM_ENABLED="true"
export TELEGRAM_TOKEN=xxx
export TELEGRAM_CHATID=yyy
export TWITTER_ENABLED="true"
export TWITTER_API_KEY=xxx
export TWITTER_API_SECRET_KEY=yyy
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_TOKEN_SECRET=yyy
```
- deploy CRDs
```
kubectl apply -f deploy/crds/community.io_weeklies_crd.yaml
kubectl apply -f deploy/crds/community.io_meetups_crd.yaml
kubectl apply -f deploy/crds/community.io_announcements_crd.yaml
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
operator-sdk build cloudnativeid/community-operator:0.0.x
```

#### Installing community-operator via helm
Please read README.md in charts folder for more information.
```
helm install ./charts --name-template community-operator --set-string telegram.chatid="xx" --set-string telegram.token="yy"
```

to insatall without crds
```
--skip-crds
```

to upgrade
```
helm upgrade community-operator ./charts --set-string telegram.chatid="xx" --set-string telegram.token="yy"
```

#### Community list
This is the community list that used community-operator
- Kubernetes & Cloud Native Indonesia (full implementation)
- Ruby Indonesia (weekly)
- Golang Indonesia (weekly)
- DevOps Indonesia (weekly)
- NodeJS Indonesia (weekly)
- JavaScript Indonesia (weekly)
