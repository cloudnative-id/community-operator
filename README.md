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

## Contributors
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://zufardhiyaulhaq.com"><img src="https://avatars3.githubusercontent.com/u/11990726?v=4" width="100px;" alt=""/><br /><sub><b>Zufar Dhiyaulhaq</b></sub></a><br /><a href="https://github.com/cloudnative-id/community-operator/commits?author=zufardhiyaulhaq" title="Code">💻</a> <a href="#tool-zufardhiyaulhaq" title="Tools">🔧</a> <a href="#ideas-zufardhiyaulhaq" title="Ideas, Planning, & Feedback">🤔</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

Contributions of any kind welcome, please check [CONTRIBUTING.md](https://github.com/cloudnative-id/community-operator/blob/master/.github/CONTRIBUTING.md)!

## Changes
For changes, see the [CHANGELOG.md](CHANGELOG.md).

## License
This program is free software: you can redistribute it and/or modify it under the terms of the [MIT license](LICENSE)
