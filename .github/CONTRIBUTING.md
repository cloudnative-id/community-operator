# Contributing
By participating to this project, you agree to abide our [code of conduct](https://github.com/cloudnative-id/community-operator/blob/master/.github/CODE_OF_CONDUCT.md).

## Development
For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:

### Setup your machine
Prerequisites:
- make
- [Go 1.13+](https://golang.org/doc/install)
- [operator-sdk v0.18.1+](https://sdk.operatorframework.io/)

Fork and clone **[community-operator](https://github.com/cloudnative-id/community-operator)** repository.

- Setup your environment
```
cp .env-example .env
set -o allexport; source .env; set +o allexport
```

- deploy CRDs
```
kubectl apply -f deploy/crds/community.io_weeklies_crd.yaml
kubectl apply -f deploy/crds/community.io_meetups_crd.yaml
kubectl apply -f deploy/crds/community.io_announcements_crd.yaml
```

- Run community-operator locally
```
operator-sdk run local
```

- deploy some examples
```
kubectl apply -f examples/weekly/example.yaml
kubectl apply -f examples/meetup/example.yaml
kubectl apply -f examples/announcement/example.yaml
```


### Submit a pull request
As you are ready with your code contribution, push your branch to your `community-operator` fork and open a pull request against the **master** branch.

Please also update the [CHANGELOG.md](https://github.com/cloudnative-id/community-operator/blob/master/CHANGELOG.md) to note what you've added or fixed.
