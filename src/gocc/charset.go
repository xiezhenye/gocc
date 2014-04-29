package gocc

type Charset [32]byte

func (self *Charset) Set(c byte, v bool) *Charset {
  i, b := pos(c)
  if v {
    self[i] |= b
  } else {
    self[i] &^= b
  }
  return self
}

func (self *Charset) Get(c byte) bool {
  i, b := pos(c)
  return (self[i] & b) != 0
}

func pos(c byte) (uint, byte) {
  var b byte
  i := c / 8;
  b = 1 << (c % 8) 
  return uint(i), b
}

func (self *Charset) SetRange(f byte, t byte, v bool) *Charset {
  if t < f {
    return self
  }
  for c := f; c <= t; c++ {
    self.Set(c, v)
  }
  return self
}

// "a-zA-Z0-9_-"
func (self *Charset) Parse(expr string) *Charset {
  bs := []byte(expr)
  var f byte
  var rg bool
  rg = false
  for _, c := range bs {
    if rg {
      self.SetRange(f + 1, c, true)
      rg = false
    } else {
      if c == '-' {
        rg = true
      } else {
        self.Set(c, true)
        f = c
      }
    }
  }
  if rg {
    self.Set('-', true)
  }
  return self
}

func (self *Charset) Empty() bool {
  for _, c := range self {
    if c != 0 {
      return false
    }
  }
  return true
}

func (self *Charset) String() string {
  b := make([]byte, 0, 256)
  for i := 0; i < 256; i++ {
    if self.Get(byte(i)) {
      b = append(b, byte(i))
    }
  }
  return string(b)
}

func (self *Charset) Clear() *Charset {
  return self.SetAll(false)
}

func (self *Charset) Full() *Charset {
  return self.SetAll(true)
}

func (self *Charset) Revert() *Charset {
  for i := range self {
    self[i] ^= '\xff'
  }
  return self
}

func (self *Charset) SetAll(v bool) *Charset {
  var c byte
  if v {
    c = '\xff'
  } else {
    c = '\x00'
  }
  for i := range self {
    self[i] = c
  }
  return self
}

func (self *Charset) IntersectWith(cs *Charset) *Charset {
  for i := range self {
    self[i] &= cs[i]
  }
  return self
}

func (self *Charset) MergeWith(cs *Charset) *Charset {
  for i := range self {
    self[i] |= cs[i]
  }
  return self
}

func (self *Charset) CopyFrom(cs *Charset) *Charset {
  copy(self[:], cs[:])
  return self
}

func CS(cs string) *Charset {
  var ret Charset
  ret.Parse(cs)
  return &ret
}

func (self *Charset) FirstSet() Charset {
  return *self
}

func (self *Charset) Nil() bool {
  return self.Empty()
}


