package db

import (
    "testing"
)

type commitInput struct {
    mod  bool       //commit mode
    cStr string     //commit string
    qStr string     //query string
}

type commitTest struct {
    in  commitInput 
    out string
}

var commitTests = []commitTest{
    commitTest{ commitInput{false, "delete from t", "select a, b from t"}, "a,b|c,d|"},
    commitTest{ commitInput{true, "delete from t", "select a, b from t"}, ""},
}

func TestCommit(t *testing.T) {
    var sql MySQL
    for _, d := range commitTests {
        v := sql.Connect("localhost", "webapi", "itbuwebapi", "webapi", 3306)
        if v!= 0 {
            t.Fatalf("Connect({localhost, webapi, itbuwebapi, webapi})=%v, want 0.", v )
        }

        v = sql.SetAutoCommit(d.in.mod)
        if v!=0 {
            t.Fatalf("SetAutoCommit(fase)=%v, want 0", v)
        }

        v = sql.Execute(d.in.cStr)
        if v!=0 {
            t.Fatalf("Execute(%v)=%v, want 0.", d.in.cStr, v)
        }

        ch := make(chan int)
        go func() {
            var sql MySQL
            v := sql.Connect("localhost", "webapi", "itbuwebapi", "webapi", 3306)
            if v!= 0 {
                t.Fatalf("Connect error.")
            }
            
            v = sql.Execute(d.in.qStr)
            if v!=0 {
                t.Fatalf("Execute(%v)=%v, want 0.", d.in.qStr, v)
            }
            var result string 
            for r,_:=sql.NextRow(); r!=nil; r,_=sql.NextRow() {
                result += r["a"]+","+r["b"]+"|"
            }
            if result != d.out {
                t.Errorf("Query(%v)=%v, want %v.", d.in.qStr, result, d.out )
            }
            ch<-0
        }()
        <-ch

        v = sql.Commit()
        if v!=0 {
            t.Fatalf("Commit=%v, want 0", v)
        }

        v = sql.Execute(d.in.qStr)
        if v!=0 {
            t.Fatalf("Execute(%v)=%v, want 0.", d.in.qStr, v)
        }

        var result string 
        for r,_:=sql.NextRow(); r!=nil; r,_=sql.NextRow() {
            result += r["a"]+","+r["b"]+"|"
        }
        if len(result)!=0 {
            t.Errorf("Query(%v)=%v, want \"\".", d.in.qStr, result)
        }
        
    }
}

