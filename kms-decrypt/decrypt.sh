#!/bin/sh

while getopts i:s: option
do
case "${option}"
in
i) AWSKEYID=${OPTARG};;
s) AWSSECRETKEY=${OPTARG};;
esac
done
shift $(expr $OPTIND - 1 )

BODY=$@

if [ ! -n "${AWSKEYID}" -o ! -n "${AWSSECRETKEY}" -o ! -n "${BODY}" ]
then
    echo Usage:
    echo - HTTP:  curl -X POST -d '-i [ AWS KEY ID ] -s [ AWS SECRET KEY ] "[ SOPS DOCUMENT CONTENTS ]"' http://127.0.0.1:31112/function/kms-decrypt
    echo - CLI:   faas invoke -g http://127.0.0.1:31112 kms-decrypt '-i [ AWS KEY ID ] -s [ AWS SECRET KEY ] "[ SOPS DOCUMENT CONTENTS ]"'
    echo
    echo ARGS:
    echo -i     \(Required\)    The AWS_ACCESS_KEY_ID to be associated with the SOPS decrytion operation
    echo -s     \(Required\)    The AWS_SECRET_ACCESS_KEY to be associated with the SOPS decryptin operation
    
else
    mkdir -p ~/.aws

    # SOPS only decrypts from a file. So, write contents to a file.
    echo -n "${BODY}" | sed 's/\\n/\n/g' > ./temp.yaml

    # Gotta have creds
    echo -n "[default]\naws_access_key_id = ${AWSKEYID}\naws_secret_access_key = ${AWSSECRETKEY}" | sed 's/\\n/\n/g' > ~/.aws/credentials

    # Do what we came here to do
    sops -d ./temp.yaml > ./temp2.yaml

    # Convert YAML output to JSON
    python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' < ./temp2.yaml

    # Remove all traces of this request
    rm ~/.aws/credentials
    rm ./temp.yaml
    rm ./temp2.yaml
fi
