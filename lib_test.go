package go_cache

import "testing"

func TestInterfaceToString(t *testing.T) {
	if ToString("t1") != "t1" {
		t.Error("ToString() string failed")
	}
	if ToString(2) != "2" {
		t.Error("ToString() int failed")
	}
	if ToString(nil) != "" {
		t.Error("ToString() nil failed")
	}

	if ToString([]byte("abc")) != "abc" {
		t.Error("ToString() bytes failed")
	}
	if ToString(1.23) != "1.23" {
		t.Error("ToString() bytes failed")
	}

	type testStruct struct {
		Name string
	}
	temp := ToString(testStruct{Name: "t1"})
	if temp != `{"Name":"t1"}` {
		t.Error("ToString() struct failed")
	}

	temp2 := map[string]any{"Name": "t1", "Age": 18}
	if ToString(temp2) != `{"Age":18,"Name":"t1"}` {
		t.Error("ToString() map failed")
	}

}
