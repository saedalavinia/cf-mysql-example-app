#!/bin/bash

set -euo pipefail

go build

cf api --skip-ssl-validation $CF_API
cf auth $CF_USERNAME $CF_PASSWORD
cf target -o $CF_ORG -s $CF_SPACE
cf create-service $CF_SERVICE_NAME $CF_SERVICE_PLAN mysql-test
cf push mysql-example-app --no-start -b https://github.com/cloudfoundry/go-buildpack.git
cf bind-service mysql-example-app mysql-test
cf start mysql-example-app

DATE=$(date)

curl --fail -k -s -d "${DATE}" -X PUT https://mysql-example-app."${CF_APP_DOMAIN}"/some-key
curl --fail -k -s https://mysql-example-app."${CF_APP_DOMAIN}"/some-key | grep -q "${DATE}"

cf delete -f mysql-example-app
cf delete-service -f mysql-test