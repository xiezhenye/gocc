package gocc

type LALR1States struct {
  states  []*LALR1State
  current uint
}

type LALR1State struct {
  act  map[Element]uint
}

