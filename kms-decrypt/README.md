# kms-decrypt

This function wraps [SOPS](https://github.com/mozilla/sops) allowing a user to take advantage SOPS decryption without needed to have it installed.

## Usage

*HTTP*
`curl -X POST -d '-i [ AWS KEY ID ] -s [ AWS SECRET KEY ] "[ SOPS DOCUMENT CONTENTS ]"' http://127.0.0.1:31112/function/kms-decrypt`

*CLI*
`faas invoke -g http://127.0.0.1:31112 kms-decrypt '-i [ AWS KEY ID ] -s [ AWS SECRET KEY ] "[ SOPS DOCUMENT CONTENTS ]"'`

## Response

JSON

*Example*
`

`

## Credentials

There are no default credentials or Kubernetes secrets integration.  The AWS Access ID and Secret Key are required to be passed with the request.