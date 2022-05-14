
**Repo Initialization**

`gcloud auth login`
- This *should* set up your system to have the default credentials setup, allowing Go to pick them up automatically

**Starting the Server**

`go run ./server.go`

**Regenerating gql generation**
`go generate ./...`

**Generation debugging**
If you get 
```
reloading module info
generating core failed: unable to load github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model - make sure you're using an import path to a package that exists
exit status 1
```

in `api/sk-serve` try running 
```
$ print '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > ./tools.go
$ echo 'package model' | gofmt > ./graph/model/doc.go
$ go get .
```