package gocc

import (
  "testing"
)

func TestGetSet(t *testing.T) {
  var cs Charset
  cs.Set('a', true)
  if !cs.Get('a') {
    t.Error("Set Get failed")
  }
  cs.Set('X', true)
  if !cs.Get('X') {
    t.Error("Set Get failed")
  }
  cs.Set(0, true)
  if !cs.Get('\000') {
    t.Error("Set Get failed")
  }
  cs.Set(255, true)
  if !cs.Get(255) {
    t.Error("Set Get failed")
  }
  cs.Set('a', false)
  if cs.Get('a') {
    t.Error("Set Get failed")
  }
}

func TestRange(t *testing.T) {
  var cs Charset
  cs.SetRange('a', 'z', true)
  if !cs.Get('a') {
    t.Error("SetRange failed")
  }
  if !cs.Get('z') {
    t.Error("SetRange failed")
  }
  if !cs.Get('k') {
    t.Error("SetRange failed")
  }
  if cs.Get('1') {
    t.Error("SetRange failed")
  }
}

func TestToString(t *testing.T) {
  var cs Charset
  cs.Set('a', true)
  cs.Set('b', true)
  s := cs.String()
  if s != "ab" {
    t.Error("String failed")
  }
  cs.SetRange('0', '9', true)
  s = cs.String()
  if s != "0123456789ab" {
    t.Error("String failed")
  }
}

func TestEmpty(t *testing.T) {
  var cs Charset
  if !cs.Empty() {
    t.Error("Empty Failed")
  }
  cs.Set('x', true)
  if cs.Empty() {
    t.Error("Empty Failed")
  }
  cs.Set('x', false) 
  if !cs.Empty() {
    t.Error("Empty Failed")
  }
}

func TestParse(t *testing.T) {
  var cs Charset
  cs.Parse("abc")
  if cs.String() != "abc" {
    t.Error("Parse failed")
  } 
  cs.Parse("0-3")
  if cs.String() != "0123abc" {
    t.Error("Parse failed")
  } 
  cs.Parse("0-5a-e_:-")
  if cs.String() != "-012345:_abcde" {
    t.Error("Parse failed")
  }
}

func TestClear(t *testing.T) {
  var cs Charset
  cs.Set('a', true)
  cs.Clear()
  if ! cs.Empty() {
    t.Error("Clear Failed")
  }
}

func TestFull(t *testing.T) {
  var cs Charset
  cs.Full()
  if ! cs.Get('x') {
    t.Error("Full Failed")
  }
}

func TestRevert(t *testing.T) {
  var cs Charset
  cs.Set('a', true).Revert()
  if cs.Get('a') {
    t.Error("Revert Failed")
  }
  if ! cs.Get('x') {
    t.Error("Revert Failed")
  }
}

func TestIntersect(t *testing.T) {
  cs1 := CS("abcde")
  cs2 := CS("abc123")
  cs1.IntersectWith(cs2)
  if cs1.String() != "abc" {
    t.Error("Intersect Failed")
  }
}

func TestMerge(t *testing.T) {
  cs1 := CS("abcde")
  cs2 := CS("abc123")
  cs1.MergeWith(cs2)
  if cs1.String() != "123abcde" {
    t.Error("Merge Failed")
  }
}

func TestCopy(t *testing.T) {
  cs1 := CS("a")
  cs2 := CS("")
  cs2.CopyFrom(cs1)
  cs1.Set('b', true)
  if cs2.Get('b') {
    t.Error("Copy Failed")
  }
  if !cs2.Get('a') {
    t.Error("Copy Failed")
  }
}

