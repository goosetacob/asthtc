# ~ a s t h t c ~
dumb tool i made because i was boread listening to mac demarco and wanted to checkout [cobraCLI](https://github.com/spf13/cobra)

## install
```bash
go get github.com/goosetacob/asthtc
```

## cli usage
#### add spaces and pre/post apend with ~
```bash
$ asthtc aesthetic goosetacob
~ g o o s e t a c o b ~

$ asthtc aesthetic los angeles
~ l o s a n g e l e s ~
```
#### remove vowels from phrase
```bash
$ asthtc voweless goosetacob
gstcb

$ asthtc voweless los angeles
lsngls
```
#### compute the [de bruijn sequence](https://en.wikipedia.org/wiki/De_Bruijn_sequence#Algorithmhttps://en.wikipedia.org/wiki/De_Bruijn_sequence#Algorithm) given order n and slphabet k
```bash
$ asthtc deBruijn --alphabet "01" --subSequenceSize 8
00010111

$ asthtc deBruijn -a "abcd" -s 2
aabacadbbcbdccdd
```

## todo
- add gRPC to communicate between CLI and backend
- add [grpc gateway](https://github.com/grpc-ecosystem/grpc-gateway) to keep RESTful JSON API
- figure out how to run backend/resource/tool_test.go
- auto dump into pbcopy (or xsel for linux) when in CLI
- figure out autocomplete for this
- something actually useful (finish implementing De Bruijn Sequence)?
