package db

import (
    "testing"
)

type prepDataTest struct {
    in  string
    out uint
}

var prepDateTests = []prepDataTest{
    prepDataTest{"drop table if exists t", 1},
    prepDataTest{"create table t (a varchar(100), b varchar(100)) ENGINE=InnoDB DEFAULT CHARSET=utf8", 1},
    prepDataTest{"insert into t(a, b) values('a', 'b')", 1},
    prepDataTest{"insert into t(a, b) values('c', 'd')", 1},
}

func TestCreateTable(t *testing.T) {
    var sql MySQL
    for _, d := range prepDateTests {
        v := sql.Connect("localhost", "webapi", "itbuwebapi", "webapi", 3306)
        if v!= 0 {
            t.Fatalf("Connect({localhost, webapi, itbuwebapi, webapi})=%v, want 0.", v )
        }

        v = sql.Execute(d.in)
        if v!=0 {
            t.Errorf("Execute(%v)=%v, want 0.", d.in, v)
        }
    }
}

