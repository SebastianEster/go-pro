/**
  (C) Copyright 2018 Armin Heller

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If ast.Not, see <https://www.gnu.org/licenses/>.
*/

package ParserCombinators

import (
  "testing"
  "github.com/jweigend/concepts-of-programming-languages/oop/ast"
  "github.com/QAhell/Parser-Gombinators/parse"
)

func TestMakeOr (t *testing.T) {
  var result = makeOr (parse.Pair { ast.Val { "a" }, ast.Val { "b" }})
  var expected ast.Node = ast.Or { ast.Val { "a" }, ast.Val { "b" } }
  if result != expected {
    t.Errorf (
      "makeOr (parse.Pair { ast.Val { \"a\" }, ast.Val { \"b\" }}) failed! Expected %v" +
      " but got wrong result %v !", expected, result)
  }
  result = makeOr (parse.Pair { ast.Val { "a" }, parse.Nothing{} })
  expected = ast.Val { "a" }
  if result != expected {
    t.Errorf (
      "makeOr (parse.Pair { ast.Val { \"a\" }, ast.Nothing{} }) failed! Expected %v " +
      " but got wrong result %v !", expected, result)
  }
}

func TestMakeAnd (t *testing.T) {
  var result = makeAnd (parse.Pair { ast.Val { "a" }, ast.Val { "b" }})
  var expected ast.Node = ast.And { ast.Val { "a" }, ast.Val { "b" } }
  if result != expected {
    t.Errorf (
      "makeAnd (parse.Pair { ast.Val { \"a\" }, ast.Val { \"b\" }}) failed! Expected %v" +
      " but got wrong result %v !", expected, result)
  }
  result = makeAnd (parse.Pair { ast.Val { "a" }, parse.Nothing{} })
  expected = ast.Val { "a" }
  if result != expected {
    t.Errorf (
      "makeAnd (parse.Pair { ast.Val { \"a\" }, ast.Nothing{} }) failed! Expected %v " +
      " but got wrong result %v !", expected, result)
  }
}

func TestMakeNot (t *testing.T) {
  var result = makeNot (0, ast.Val { "a" })
  var expected ast.Node = ast.Val { "a" }
  if result != expected {
    t.Errorf ("makeast.Not (0, ast.Val { \"a\" }) failed! Expected %v " +
      "but got wrong result %v !", expected, result)
  }
  expected = ast.Not { ast.Not { ast.Not { ast.Val { "a" } }}}
  result = makeNot (3, ast.Val { "a" })
  if result != expected {
    t.Errorf ("makeast.Not (3, ast.Val { \"a\" }) failed! Expected %v " +
      "but got wrong result %v !", expected, result)
  }
}

func TestParseVariable (t *testing.T) {
  var text = "xyz"
  var expected ast.Node = ast.Val { "xyz" }
  var result = parseVariable (parse.StringToInput (text))
  if result.Result != expected {
    t.Errorf ("parseVariable on input \"%v\" failed! Expected %v " +
      "but got wrong result %v !", text, expected, result.Result)
  }
  if result.RemainingInput != nil {
    var inp  = result.RemainingInput.(parse.RuneArrayInput)
    var rest = inp.Text[inp.CurrentPosition:]
    t.Errorf ("parseVariable didn't eat all the input. Leftover: \"%v\"",
      string (rest))
  }
}

func TestParseExclamationMarks (t *testing.T) {
  var text = "!!!x"
  var expected int = 3
  var result = parseExclamationMarks (parse.StringToInput (text))
  if result.Result != expected {
    t.Errorf ("parseExclamationMarks on input \"%v\" failed! Expected %d " +
      "but got wrong result %d !", text, expected, result.Result)
  }
  if result.RemainingInput != nil {
    var inp  = result.RemainingInput.(parse.RuneArrayInput)
    var rest = inp.Text[inp.CurrentPosition:]
    if ("x" != string (rest)) {
      t.Errorf ("parseExclamationMarks ate the wrong amout of input! " +
        "Leftover: \"%v\"", string (rest))
    }
  } else {
    t.Errorf ("parseExclamationMarks mustn't eat all the input of \"%v\" " +
      "but it did!", text)
  }

}

func testExp (t *testing.T, text string, expected ast.Node) {
  var result = parseExpression(parse.StringToInput (text))
  if result.Result != expected {
    t.Errorf ("parseExpression on input \"%v\" failed! Expected %v " +
      "but got wrong result %v !", text, expected, result.Result)
  }
  if result.RemainingInput != nil {
    var inp  = result.RemainingInput.(parse.RuneArrayInput)
    var rest = inp.Text[inp.CurrentPosition:]
    t.Errorf ("parseExpression didn't eat all the input. " +
      "Leftover: \"%v\"", string (rest))
  }
}

func TestParseExpression (t *testing.T) {
  testExp (t, "!a", ast.Not { ast.Val { "a" } })
  testExp (t, "a&b", ast.And { ast.Val {"a"}, ast.Val {"b"} })
  testExp (t, "a|b", ast.Or { ast.Val {"a"}, ast.Val {"b"} })
  testExp (t, " a &  b", ast.And { ast.Val {"a"}, ast.Val {"b"} })
  testExp (t, "a   ", ast.Val {"a"})
  testExp (t, "a&b&c", ast.And { ast.Val{"a"}, ast.And { ast.Val{"b"}, ast.Val{"c"} }})
  testExp (t, "a&(b&c)", ast.And { ast.Val{"a"}, ast.And { ast.Val{"b"}, ast.Val{"c"} }})
  testExp (t, "(a&b)&c", ast.And { ast.And { ast.Val{"a"}, ast.Val{"b"} }, ast.Val{"c"}})
  testExp (t, "a|b|c", ast.Or { ast.Val{"a"}, ast.Or { ast.Val{"b"}, ast.Val{"c"} }})
  testExp (t, "a|(b|c)", ast.Or { ast.Val{"a"}, ast.Or { ast.Val{"b"}, ast.Val{"c"} }})
  testExp (t, "(a|b)|c", ast.Or { ast.Or { ast.Val{"a"}, ast.Val{"b"} }, ast.Val{"c"}})
  testExp (t, "!a & b|c&!(d|e)",
    ast.Or { ast.And { ast.Not { ast.Val { "a" } }, ast.Val { "b" } },
            ast.And { ast.Val { "c" }, ast.Not {ast.Or { ast.Val { "d" }, ast.Val { "e" } }}}})

}
