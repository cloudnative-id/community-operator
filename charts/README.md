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
| operator.tag | Tag for image community-operator | 0.0.4 |
| operator.pullPolicy | pullPolicy | Always |
| operator.replica | number of replica | 3 |
| telegram.enabled | Telegram enabled | true |
| telegram.chatid | Telegram chatid | |
| telegram.token | Telegram bot token | |
| twitter.enabled | Telegram enabled | true |
| twitter.api_key | Twitter API key | |
| twitter.api_secret_key | Twitter API secret key | |
| twitter.access_token | Twitter access token | |
| twitter.access_token_secret | TTwitter access token secrets | |

Specify each parameter using the `--set key=value[,key=value]` argument to helm install. For example:
```
helm install ./charts --name-template community-operator --set-string telegram.chatid="-1234556" --set-string telegram.token="12354:asdaADASFD"
```
