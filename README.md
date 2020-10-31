# community-operator
The Community Operator provides Kubernetes native deployment and management of your community. The purpose of this project is to simplify and automate community management in top of Kubernetes clusters.

Community-operator currently watch CRDs on the same namespace as the community-operator deployed.

## Feature
- **Kubernetes Object based**
- **Support Weekly, Announcement, and Meetup notification**
- **Minimal configuration**

## Installation

### Helm
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

## Users
This is the community list that used community-operator
- Kubernetes & Cloud Native Indonesia 
- Ruby Indonesia
- Golang Indonesia
- DevOps Indonesia
- NodeJS Indonesia
- JavaScript Indonesia
- OpenStack Indonesia
- Open Networking Indonesia

## Changes
For changes, see the [CHANGELOG.md](CHANGELOG.md).

## License
This program is free software: you can redistribute it and/or modify it under the terms of the [MIT license](LICENSE)
