package coverage

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
type testCases struct{
	descr string
	input People
	want interface{}
}

type testCaseMatrix struct{
	descr string
	input *Matrix
	matrixStr string
}

type testCaseRowsCols struct{
	descr string
	input [][]int
	matrix string
}

func TestLen(t *testing.T){

	t.Parallel()

	p := Person{}

	for _,tcase := range []testCases{
		{
			descr: "1 person",
			input: People{p},
			want:1,

		},
		{
			descr: "4 persons",
			input: People{p,p,p,p},
			want: 4,
		},
	}{
		t.Run(tcase.descr,func(t *testing.T){
			got := tcase.input.Len()
			if got != tcase.want{
				t.Errorf("Expected %v, but got %v",tcase.want,got)
			}
		})
	}
}

func TestLess(t *testing.T){

	t.Parallel()

	//Test data
	p1 := Person{"John","Down",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p2 := Person{"John","Grayhouse",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p3 := Person{"Gregory","Downtown",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p4 := Person{"Antony","Woodpecker",time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)}

	for _,tcase := range []testCases{
		{
			descr: "First is older",
			input: People{p1,p4},
			want: true,
		},
		{
			descr: "First is younger",
			input: People{p4,p1},
			want: false,
		},
		{
			descr: "First is older,name is smaller",
			input: People{p1,p3},
			want: true,
		},
		{
			descr: "First is younger,name is bigger",
			input: People{p3,p1},
			want: false,
		},
		{
			descr: "Same age, same name, last name is smaller",
			input: People{p1,p2},
			want: false,
		},
		{
			descr: "Same age, same name, last name is bigger",
			input: People{p2,p1},
			want: true,
		},
	}{
		t.Run(tcase.descr,func(t *testing.T) {
			got := tcase.input.Less(1,0)
			if got != tcase.want{
				t.Errorf("Expected %t, but got %t",tcase.want,got)
			}
		})
	}
}


func TestSwap(t *testing.T){
	t.Parallel()
	
	// Test data
	p1:= Person{"John","Dow",time.Now()}
	p2:= Person{"Greg","Frank",time.Now()}

	testData := People{p1,p2}

	testData.Swap(0,1)

	if testData[0] != p2{
		t.Errorf("Person 2 expected to be first element.")
	}

}

func TestNewMatrix(t *testing.T){
	t.Parallel()

	for _,tcase := range []testCaseMatrix{
		{
			descr:"basic matrix 3x3",
			input:&Matrix{3,3,[]int{5,6,7,7,7,7,7,6,7}},
			matrixStr:"5 6 7\n7 7 7\n7 6 7",
		},
		{
			descr:"matrix 2x2",
			input: &Matrix{2,2,[]int{4,6,6,7}},
			matrixStr:"4 6\n6 7",
		},
		{
			descr:"matrix with spaces",
			input: &Matrix{2,2,[]int{5,6,6,1}},
			matrixStr:" 5 6 \n 6 1",
		},
		
	}{
		t.Run(tcase.descr,func(t *testing.T) {
			got,_ := New(tcase.matrixStr)
			assert.Equal(t,got.data,tcase.input.data)
		})
	}
}

func TestNewMatrixWithErrors(t *testing.T){
	t.Parallel()

	for _, tcase := range []testCaseMatrix{
		{
			descr:"rows>cols",
			input: &Matrix{2,3,[]int{5,5,5}},
			matrixStr:"5 4\n6 4 5",
		},
		{
			descr:"wrong format of values",
			input: &Matrix{3,2,[]int{4,5,6}},
			matrixStr:"6 h 5\n7 7 7",
		},
		{
			descr:"cols>rows",
			input: &Matrix{3,2,[]int{5}},
			matrixStr:"4 5 6\n5 6",
		},
		{
			descr:"empty matrix",
			input: &Matrix{0,0,[]int{}},
			matrixStr:"",
		},
	}{
		t.Run(tcase.descr,func(t *testing.T) {
			_ ,err := New(tcase.matrixStr)
			assert.NotNil(t,err)
		})
	}
}


func TestMatrixRows(t *testing.T){
	t.Parallel()

	for _, tcase := range []testCaseRowsCols{

		{
			descr: "1 row",
			input: [][]int{{4,5,4}},
			matrix: "4 5 4",
		},
		{
			descr: "multiple rows",
			input: [][]int{{7, 5, 5}, {6, 4, 5}, {9, 9, 9}},
			matrix: "7 5 5\n6 4 5\n9 9 9",
		},
	}{
		t.Run(tcase.descr,func(t *testing.T) {
			m, _ := New(tcase.matrix)
			assert.Equal(t,m.Rows(),tcase.input)
		})
	}
}

func TestMatrixCols(t *testing.T){
	t.Parallel()

	for _, tcase := range []testCaseRowsCols{

		{
			descr: "1 col",
			input: [][]int{{4},{5},{4}},
			matrix: "4 5 4",
		},
		{
			descr: "multiple cols",
			input: [][]int{{7,6,9},{5,4,9},{5,5,9}},
			matrix: "7 5 5\n6 4 5\n9 9 9",
		},
	}{
		t.Run(tcase.descr,func(t *testing.T) {
			m, _ := New(tcase.matrix)
			assert.Equal(t,m.Cols(),tcase.input)
		})
	}
}


func TestMatrixSet(t *testing.T){

	t.Parallel()

	type test struct{
		name string
		input string
		row,col,value int
	}

	testCases := []test{
		{"1 example","5 5\n6 1",1,1,0},
		{"2 example","5 5 7\n6 5 7",1,2,0},
	}

	errorCases := []test{
		{"negative row","5 5\n6 1",-1,1,0},
		{"row out of range","5 5 7\n6 5 7",3,2,0},
		{"negative col","5 5 7\n6 5 7",2,-2,0},
		{"col out of range","5 5 7\n6 5 7",1,5,0},
	}

	for _,testCase := range testCases{
		m,_ := New(testCase.input)
		assert.True(t,m.Set(testCase.row,testCase.col,testCase.value))
	}

	for _,errCase := range errorCases{

		m,_ := New(errCase.input)

		if m.Set(errCase.row,errCase.col,errCase.value){
			t.Errorf("Should be error in case %s.",errCase.name)
		}
	}
}

