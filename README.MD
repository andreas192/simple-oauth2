# Coding Exercise: Simple OAuth2 Server

## Tasks

*    Create a Golang http server that issues JWT Access Tokens (rfc7519) using Client Credentials Grant with Basic Authentication (rfc6749) in endpoint /token
*    Sign the tokens with a self-made RS256 key
*    Provide an endpoint to list the signing keys (rfc7517)
*    Provide deployment manifests to deploy the server in Kubernetes cluster
*    Create an Introspection endpoint (rfc7662) to introspect the issued JWT Access Tokens

## Remarks

* Publish the exercise in a git server and grant us access to review it.
* Avoid a single commit with the whole solution, we want to see how the solution was developed incrementally.
* Provide instructions to execute it


## Execute

*   Clone repo
*   Use https://keytool.online/ to generate PEM(PKCS#1) public and private key and populate public-key.pem and private-key.pem
*   Run in terminal `go mod download`
*   Run in terminal go run .
*   Test JWT Generate API:
curl --location 'localhost:8090/token' \
--header 'Content-Type: application/json' \
--data '{
    "username": "<your user>",
    "password": "<your password>"
}'
*   Test Signing Key API
curl --location 'localhost:8090/signing-keys' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODY2ODQxODYsImlzcyI6InlvdXItaXNzdWVyIiwic3ViIjoiY2xpZW50In0.MwMVoA5ttjxG12b8c0EbTuki5t2vwnZTjlGPZtfJFm3-SFGuPyAj7pK9weilbVULbZ4-6IoFdtvEQT4PQZ3uNpWw-zjClI-CgIfvbs_D97xUaj_a48CfOrYzUCFgDK_xesfH5tXcKLOBn3O5Yg4kR4BTQEff-Z9XFH0B6HLQP2ZzGFnBLuNy89F_FG6lVfsnaj4agscigSOCTvn61T8GXnF8pZc2fbdUPiDlpdCFq7jbBwRT0fQ55jvpu2QhDWOQdEFPWob8ll6yQVSVk0md4FQCrwRJ7C1Kk7p2Aw5NmHDVaqzG1wtdBUtMWVm-tGZOA1LaAuR-0SXpaWebIOsSzw'
*   Test Introspection API
curl --location 'localhost:8090/introspection' \
--header 'Authorization: Bearer <access_token from /token API>'
*   Celebrate!!!

