package gocc

import "testing"

func TestEBNF(t *testing.T) {
  ignore := EBNF{}
  ignore.
    Add("SP", Rep(CS("\x00-\x20")))

  rules := EBNF{}
  rules.
    Add("name",    CS("a-zA-Z_"), Rep(CS("0-9a-zA-Z_:"))).
    Add("string",  Ch('"'), Rep( CS("\"").Revert() ), Ch('"')).
    Add("attr",    rules.Rule("name"), Ch('='), rules.Rule("string")).
    Add("attrs",   Ch('('), Rep(rules.Rule("attr")), Ch(')'))
/*
    Add("subs",    Ch('{'), Rep(rules.Rule("node")), Ch('}')).
    Add("value",   rules.Rule("string")).
    Add("body",    Alt({ rules.Rule("subs") }, { rules.Rule("value") })*/
  print( rules)
}


