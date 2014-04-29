package gocc

type EBNF map[string] *Rule

type Rule []Element

type Element interface {
  FirstSet() Charset
  Nil() bool
}

type Str string

type Ch byte

func (self *EBNF) Add(name string, e ... Element) *EBNF {
  (*self)[name] = (*Rule)(&e)
  return self
}

func Rep(e Element) Element {
  return nil
}

func Opt(e Element) Element {
  return nil
}

func Grp(e Element) Element {
  return nil
}

func Alt(e Element) Element {
  return nil
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

func (self *Rule) FirstSet() Charset {
  var ret Charset
  for _, e := range *self {
    fs := e.FirstSet()
    ret.MergeWith(&fs)
    if ! e.Nil() {
      break
    }
  }
  return ret
}

func (self *Rule) Nil() bool {
  for _, e := range *self {
    if ! e.Nil() {
      return false
    }
  }
  return true
}

