# community-operator charts
Helm chart for community-operators

### Installing the charts
From root directory of community-operator. Please edit the values.yaml inside charts before applying.
```
helm install ./charts --name community-operator
```

### Configuration

| Parameter | Description | Default |
|-|-| -|
| operator.image | Image for community-operator | cloudnativeid/community-operator |
| operator.tag | Tag for image community-operator | 0.0.1 |
| operator.pullPolicy | pullPolicy | Always |
| operator.replica | number of replica | 3 |
| telegram.chatid | Telegram chatid | |
| telegram.token | Telegram bot token | |

Specify each parameter using the `--set key=value[,key=value]` argument to helm install. For example:
```
helm install ./charts --name-template community-operator --set-string telegram.chatid="-1234556" --set-string telegram.token="12354:asdaADASFD"
```
