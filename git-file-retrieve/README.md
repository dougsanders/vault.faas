# git-file-retrieve

This function will log into a git repository and return the contents of files that match the given regex.

## Usage

*HTTP*

`curl -X POST -d '{"repoUrlHttp":"https://github.com/somerepo.git", "pathRegex":"somepath\/.*"}' http://127.0.0.1:31112/function/git-file-retrieve`

*CLI*

`faas invoke -g http://127.0.0.1:31112 git-file-retrieve '{"repoUrlHttp":"https://github.com/somerepo.git", "pathRegex":"somepath\/.*"}'`

## Response

JSON

```json
{
    "files": [
        {
            "file": "dir/subdir/example.yaml",
            "contents": "Contents of example.yaml"
        }
    ]
}
```

## Credentials

This function expects there to be a generic Kubernetes secret to be setup named "ci-github-credentials".  The secret must contain values for "username" and "password"

`kubectl create secret generic ci-github-credentials -n openfaas-fn`
`kubectl edit secret ci-github-credentials -n openfaas-fn`

*Before*

```
apiVersion: v1
kind: Secret
metadata:
  creationTimestamp: "2019-01-01T00:00:00Z"
  name: example
  namespace: openfaas-fn
  resourceVersion: "100000"
  selfLink: /api/v1/namespaces/openfaas-fn/secrets/example
  uid: 54b40cf4-1da2-11e9-bdc5-025000000001
type: Opaque
```

*After*

```
apiVersion: v1
data:
  password: << ADD BASE64 ENCODED VALUE HERE >>
  username: << ADD BASE64 ENCODED VALUE HERE >>
kind: Secret
metadata:
  creationTimestamp: "2019-01-01T00:00:00Z"
  name: example
  namespace: openfaas-fn
  resourceVersion: "100000"
  selfLink: /api/v1/namespaces/openfaas-fn/secrets/example
  uid: 54b40cf4-1da2-11e9-bdc5-025000000001
type: Opaque
```

