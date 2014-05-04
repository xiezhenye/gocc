package gocc

type EBNF map[string] []Element

type Element interface {
  FirstSet() Charset
  Nil() bool
}

type Rule struct {
  Name    string
  EBNF    *EBNF
}

type Elements []Element

func (self Elements) FirstSet() Charset {
  var ret Charset
  for _, e := range self {
    fs := e.FirstSet()
    ret.MergeWith(&fs)
    if ! e.Nil() {
      break
    }
  }
  return ret
}

func (self Elements) Nil() bool {
  for _, e := range self {
    if ! e.Nil() {
      return false
    }
  }
  return true
}

type Str string

type Ch byte

func (self *EBNF) Add(name string, eles ... Element) *EBNF {
  (*self)[name] = eles 
  return self
}

func (self *EBNF) Rule(name string) Rule {
  return Rule{ Name: name, EBNF: self }
}

type Repetition struct {
  Elements []Element
}

type Option struct {
  Elements []Element
}

type Alternation struct {
  Elements []Elements
}

func (self *Repetition) FirstSet() Charset {
  return Elements(self.Elements).FirstSet()
}

func (self *Repetition) Nil() bool {
  return true
}

func (self *Option) FirstSet() Charset {
  return Elements(self.Elements).FirstSet()
}

func (self *Option) Nil() bool {
  return true
}

func (self *Alternation) FirstSet() Charset {
  var ret Charset
  for _, e := range self.Elements {
    fs := Elements(e).FirstSet()
    ret.MergeWith(&fs)
    if ! e.Nil() {
      break
    }
  }
  return ret
}

func (self *Alternation) Nil() bool {
  for _, e := range self.Elements {
    if ! Elements(e).Nil() {
      return false
    }
  }
  return true
}

func Rep(e ...Element) Element {
  return &Repetition{ e }
}

func Opt(e ...Element) Element {
  return &Option{ e }
}

func Alt(e ...Elements) Element {
  return &Alternation{ e }
}

func (self Str) FirstSet() Charset {
  var ret Charset
  ret.Set(self[0], true)
  return ret
}

func (self Str) Nil() bool {
  return len(self) > 0
}

func (self Ch) FirstSet() Charset {
  var ret Charset
  ret.Set(byte(self), true)
  return ret
}

func (self Ch) Nil() bool {
  return false
}

func (self Rule) FirstSet() Charset {
  return Elements((*self.EBNF)[self.Name]).FirstSet()
}

func (self Rule) Nil() bool {
  return Elements((*self.EBNF)[self.Name]).Nil()
}

