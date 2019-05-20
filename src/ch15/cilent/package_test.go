package cilent

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	// t.Log(series.square(6))  //小写的不能被访问到
}
