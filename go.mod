module github.com/cloudnative-id/community-operator

go 1.13

require (
	github.com/circonus-labs/circonus-gometrics v2.3.1+incompatible
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gomodule/oauth1 v0.0.0-20181215000758-9a59ed3b0a84
	github.com/norwoodj/helm-docs v1.4.0 // indirect
	github.com/operator-framework/operator-sdk v0.18.1
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/pflag v1.0.5
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/zufardhiyaulhaq/ngrok-operator v0.0.0-20201031134108-a79211c21da1 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.18.2 // Required by prometheus-operator
)
