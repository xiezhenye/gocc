package gocc

import "testing"

func TestEBNF(t *testing.T) {
  ignore := EBNF{}
  ignore.
    Add("blank", Rep(CS("\x00-\x20")))

  rules := EBNF{}
  rules.
    Add("name",    CS("a-zA-Z_"), Rep(CS("0-9a-zA-Z_:"))).
    Add("string",  Ch('"'), Rep( &Rule{ CS("\"").Revert() } ), Ch('"')).
    Add("attr",    rules["name"], Ch('='), rules["string"]).
    Add("attrs",   Ch('('), Rep(rules["attr"]), Ch(')'))
  print( rules)
}


