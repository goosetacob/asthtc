# ~ a s t h t c ~
dumb tool i made because i was boread listening to mac demarco and wanted to checkout [cobraCLI](https://github.com/spf13/cobra)

## get all the stuff
```bash
$ go get -d github.com/goosetacob/asthtc
```

## backend
### setup
```bash
# go to the code
$ cd $GOPATH/src/github.com/goosetacob/asthtc

# (optional if go source in proto/ is bad) build api/tool.pb.go
$ make grpc-go-source

# build docker images
$ make images

# run backend container
$ make backend-run
```

## asthtcCLI
### install
```bash
$ go install ./asthtcCLI
```
### usage
#### add spaces and pre/post apend with ~
```bash
$ asthtcCLI aesthetic --phrase 'goosetacob'
~ g o o s e t a c o b ~

$ asthtcCLI aesthetic --phrase 'los angeles'
~ l o s a n g e l e s ~
```
#### remove vowels from phrase
```bash
$ asthtcCLI voweless --phrase 'goosetacob'
gstcb

$ asthtcCLI voweless --phrase 'los angeles'
lsngls
```
#### compute the [de bruijn sequence](https://en.wikipedia.org/wiki/De_Bruijn_sequence#Algorithmhttps://en.wikipedia.org/wiki/De_Bruijn_sequence#Algorithm) given order n and slphabet k
```bash
$ asthtcCLI --alphabet "01" --subSequenceSize 8
00010111

$ asthtcCLI -a "abcd" -s 2
aabacadbbcbdccdd
```

## JSON over HTTP
### setup
need reverse-proxy to transtalte JSON/HTTP to RPC/HTTP
```bash
$ make reverse-proxy-run
```
### usage
#### remove vowels from phrase
```bash
$ curl -d '{"Phrase":"los angeles"}' -H "Content-Type: application/json" -XPOST http://localhost:8080/v1/voweless
gstcb

$ curl -d '{"Phrase":"goosetacob"}' -H "Content-Type: application/json" -XPOST http://localhost:8080/v1/voweless
{"Phrase":"gstcb"}
```

## todo
- ~~add gRPC to communicate between CLI and backend~~
- ~~add [grpc gateway](https://github.com/grpc-ecosystem/grpc-gateway) to keep RESTful JSON API as an option~~
    - figure out why [tools.swagger.json](proto/toolsService/tools.swagger.json) looks wrong
- figure out how to run backend/resource/tool_test.go
- auto dump into pbcopy (or xsel for linux) when in CLI
- figure out autocomplete for this
- something actually useful (finish implementing De Bruijn Sequence)?
- probably a few memory leaks around connections in CLI, figure out how to close those
