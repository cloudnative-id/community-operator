# community-operator charts
Helm chart for community-operators

### Installing the charts
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/community-operator --name-template community-operator
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| config.telegram.channel[0].token | string | `"bot_token"` |  |
| config.telegram.channel[0].username | string | `"username_channel_one"` |  |
| config.telegram.channel[1].token | string | `"bot_token"` |  |
| config.telegram.channel[1].username | string | `"username_channel_two"` |  |
| config.telegram.enabled | bool | `true` |  |
| config.telegram.group[0].chatId | string | `"group_one_chatid"` |  |
| config.telegram.group[0].token | string | `"bot_token"` |  |
| config.telegram.group[1].chatId | string | `"group_two_chatid"` |  |
| config.telegram.group[1].token | string | `"bot_token"` |  |
| config.twitter.account[0].accessToken | string | `"account_0_access_token"` |  |
| config.twitter.account[0].accessTokenSecret | string | `"account_0_access_token_secret"` |  |
| config.twitter.account[0].apiKey | string | `"account_0_api_key"` |  |
| config.twitter.account[0].apiSecretKey | string | `"account_0_api_secret_key"` |  |
| config.twitter.account[1].accessToken | string | `"account_1_access_token"` |  |
| config.twitter.account[1].accessTokenSecret | string | `"account_1_access_token_secret"` |  |
| config.twitter.account[1].apiKey | string | `"account_1_api_key"` |  |
| config.twitter.account[1].apiSecretKey | string | `"account_1_api_secret_key"` |  |
| config.twitter.enabled | bool | `false` |  |
| operator.image | string | `"cloudnativeid/community-operator"` |  |
| operator.pullPolicy | string | `"Always"` |  |
| operator.replica | int | `1` |  |
| operator.tag | string | `"0.0.6"` |  |

Specify each parameter using the `--set key=value[,key=value]` argument to helm install. For example:
```
helm install zufardhiyaulhaq/community-operator --name-template community-operator --set-string telegram.chatid="-1234556" --set-string telegram.token="12354:asdaADASFD"
```
