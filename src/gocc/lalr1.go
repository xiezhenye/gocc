package gocc

type LALR1States struct {
  states  []*State
  current uint
}

type LALR1State struct {
  map  map[Element]uint
}

