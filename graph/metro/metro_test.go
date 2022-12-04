package main

import "testing"

/*
															   3 min
													taganskaya	---	 kitay gorod
														|				|
														|5 min			|3 min
			  3 min		  3 min			    3 min		|	  3 min		|
	novogireevo --- perovo --- aviamotornaya --- marksistskaya --- tretyakovskaya
*/

func TestMetro(t *testing.T) {
	metro := NewMoscowMetro()

	a := metro.CalculateRoute(PEROVO, KITAY_GOROD)
	b := metro.CalculateRoute(PEROVO, NOVOGIREEVO)
	c := metro.CalculateRoute(PEROVO, KITAY_GOROD)
	d := metro.CalculateRoute(TAGANSKAYA, KITAY_GOROD)
	e := metro.CalculateRoute(NOVOGIREEVO, TAGANSKAYA)

	expected := [...]bool{true, true, true, true, true}
	actual := [...]bool{a, b, c, d, e}

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
