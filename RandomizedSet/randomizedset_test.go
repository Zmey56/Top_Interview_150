package RandomizedSet

import (
	"math/rand"
	"testing"
	"time"
)

func Test_RandomizedSet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Run("Test Insert", func(t *testing.T) {
		rs := Constructor()
		if rs.Insert(1) != true {
			t.Errorf("Insert(1) failed")
		}
		if rs.Insert(1) != false {
			t.Errorf("Insert(1) failed")
		}
		if rs.Insert(2) != true {
			t.Errorf("Insert(2) failed")
		}
		if rs.Insert(2) != false {
			t.Errorf("Insert(2) failed")
		}
	})

	t.Run("Test Remove", func(t *testing.T) {
		rs := Constructor()
		if rs.Remove(1) != false {
			t.Errorf("Remove(1) failed")
		}
		rs.Insert(1)
		if rs.Remove(1) != true {
			t.Errorf("Remove(1) failed")
		}
		if rs.Remove(1) != false {
			t.Errorf("Remove(1) failed")
		}
	})

	t.Run("Test GetRandom", func(t *testing.T) {
		rs := Constructor()
		rs.Insert(1)
		rs.Insert(2)
		rs.Insert(3)
		randomVal := rs.GetRandom()
		if randomVal != 1 && randomVal != 2 && randomVal != 3 {
			t.Errorf("GetRandom() failed")
		}
	})
}

func BenchmarkConstructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor()
	}
}

func BenchmarkInsert(b *testing.B) {
	rs := Constructor()
	for i := 0; i < b.N; i++ {
		rs.Insert(i)
	}
}

func BenchmarkRemove(b *testing.B) {
	rs := Constructor()
	for i := 0; i < b.N; i++ {
		rs.Remove(i)
	}
}

func BenchmarkGetRandom(b *testing.B) {
	rs := Constructor()
	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)
	for i := 0; i < b.N; i++ {
		rs.GetRandom()
	}
}

func BenchmarkRandomizedSet(b *testing.B) {
	rs := Constructor()
	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)
	rs.Insert(4)
	rs.Insert(5)
	rs.Insert(6)
	rs.Insert(7)
	rs.Insert(8)
	rs.Insert(9)
	rs.Insert(10)
	for i := 0; i < b.N; i++ {
		rs.GetRandom()
	}
	rs.Remove(7)
}

func BenchmarkRandomizedSetVer2(b *testing.B) {
	rs := ConstructorVer2()
	rs.InsertVer2(1)
	rs.InsertVer2(2)
	rs.InsertVer2(3)
	rs.InsertVer2(4)
	rs.InsertVer2(5)
	rs.InsertVer2(6)
	rs.InsertVer2(7)
	rs.InsertVer2(8)
	rs.InsertVer2(9)
	rs.InsertVer2(10)
	for i := 0; i < b.N; i++ {
		rs.GetRandomVer2()
	}
	rs.RemoveVer2(7)
}

func BenchmarkRandomizedSetVer3(b *testing.B) {
	rs := ConstructorVer3()
	rs.InsertVer3(1)
	rs.InsertVer3(2)
	rs.InsertVer3(3)
	rs.InsertVer3(4)
	rs.InsertVer3(5)
	rs.InsertVer3(6)
	rs.InsertVer3(7)
	rs.InsertVer3(8)
	rs.InsertVer3(9)
	rs.InsertVer3(10)
	for i := 0; i < b.N; i++ {
		rs.GetRandomVer3()
	}
	rs.RemoveVer3(7)
}
