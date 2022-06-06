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

func TestLen(t *testing.T){


	p := Person{"John","Dow",time.Now()}

	
	type test struct{
		input People
		want int
	}

	tests := []test{
		{input: People{p,p,p},want: 3},
		{input: People{p,p,p,p},want: 4},
		{input: People{},want: 0},
	}


	for _,test := range tests{
		got := test.input.Len()

		if got != test.want{
			t.Errorf("Expected want %d , but got %d",test.want,got)
		}

	}
}

func TestLess(t *testing.T){

	type test struct{
		input People
		want bool
	}


	p1 := Person{"John","Down",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p2 := Person{"John","Grayhouse",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p3 := Person{"Gregory","Downtown",time.Date(1999, 8, 2, 0, 0, 0, 0, time.UTC)}
	p4 := Person{"Antony","Woodpecker",time.Date(2014, 2, 1, 0, 0, 0, 0, time.UTC)}

	tests := []test{
		{input: People{p1,p4},want: true},
		{input: People{p4,p1},want: false},
		{input: People{p1,p3},want: true},
		{input: People{p3,p1},want: false},
		{input: People{p1,p2},want: false},
		{input: People{p2,p1},want: true},
	}

	for _,test := range tests{
		

		if test.input.Less(1,0) != test.want{
			t.Errorf("Expected %t, but got %t",test.want,test.input.Less(1,0))
		}
		

	}

}


func TestSwap(t *testing.T){

	p1:= Person{"John","Dow",time.Now()}
	p2:= Person{"Greg","Frank",time.Now()}

	testData := People{p1,p2}

	testData.Swap(0,1)

	if testData[0] != p2{
		t.Errorf("Person 2 expected to be first element.")
	}


}


func TestNewMatrix(t *testing.T){

	type test struct{
		name string
		input string
		cols,rows int
		data []int
	}

	validCases := []test{
		{"basic matrix 3x3","5 6 7\n 7 7 7",3,3,[]int{5,6,7,7,7,7}},
		{"matrix 2x2","4 6\n6 7",2,2,[]int{4,6,6,7}},
		{"matrix with spaces"," 5 6 \n 6 1",2,2,[]int{5,6,6,1}},

		
	}
	errorCases :=[]test{
		{"rows>cols","5 4\n6 4 5",2,3,[]int{5,5,5}},
		{"wrong format of values","6 h 5\n7 7 7",3,2,[]int{4,5,6}},
		{"cols>rows","4 5 6\n5 6",3,3,[]int{5}},
		{"empty matrix","",0,0,[]int{6,6,4}},

	}

	for _,validTest := range validCases{
		m,_ := New(validTest.input)

		if len(m.data) != len(validTest.data){
			t.Errorf("Error in case %s",validTest.name)
		}	
	}

	for _,ererrorCases := range errorCases{
		_,err := New(ererrorCases.input)
		if err != nil{
			continue
		}
	}
}


func TestMatrixRows(t *testing.T){

	type test struct{
		name string
		input string
		expectedRows [][]int
	}

	testCases := []test{
		{"1 row","5 7 1",[][]int{{5,7,1}}},
		{"multiple rows","5 7 1\n7 8 6",[][]int{{5,7,1},{7,8,6}}},
	}

	for _,testCase := range testCases{

		m,_ := New(testCase.input)
		assert.Equal(t,testCase.expectedRows,m.Rows())
	}
}

func TestMatrixCols(t *testing.T){

	type test struct{
		name string
		input string
		expectedRows [][]int
	}

	testCases := []test{
		{"1 col","5 7 1",[][]int{{5},{7},{1}}},
		{"multiple col","5 7 1\n7 8 6",[][]int{{5,7},{7,8},{1,6}}},
	}

	for _,testCase := range testCases{

		m,_ := New(testCase.input)
		assert.Equal(t,testCase.expectedRows,m.Cols())
	}
}


func TestMatrixSet(t *testing.T){

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