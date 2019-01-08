# vault.faas

echo -n $PASSWORD | faas-cli login -g http://$OPENFAAS_URL -u admin --password-stdin

faas-cli list -g https://faas.awsdev.onetechnologies.net


faas new --lang golang-http faas-vault-renew
go get -u github.com/golang/dep/cmd/dep
dep ensure -add github.com/openfaas-incubator/go-function-sdk

faas build -f ./stacks/aws.development.vault.renew.yml
faas push -f ./stacks/aws.development.vault.renew.yml
faas deploy -f ./stacks/aws.development.vault.renew.yml
