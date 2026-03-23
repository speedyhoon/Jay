package main

import "time"

type One struct {
	One bool
}

type Two struct {
	One
	Two []bool
}

type Three struct {
	//One
	Two
	Three byte
}

type Four struct {
	//One
	//Two
	Three
	Four []byte
}

type Five struct {
	//One
	//Two
	//Three
	Four
	Five complex64
}

type Six struct {
	//One
	//Two
	//Three
	//Four
	Five
	Six []complex64
}

type Seven struct {
	//One
	//Two
	//Three
	//Four
	//Five
	Six
	Seven complex128
}

type Eight struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	Seven
	Eight []complex128
}

type Nine struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	Eight
	Nine float32
}

type Ten struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	Nine
	Ten []float32
}

type Eleven struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	Ten
	Eleven float64
}

type Twelve struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	Eleven
	Twelve []float64
}

type Thirteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	Twelve
	Thirteen int
}

type Fourteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	Thirteen
	Fourteen int8
}

type Fifteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	Fourteen
	Fifteen int16
}

type Sixteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	Fifteen
	Sixteen int32
}

type Seventeen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	Sixteen
	Seventeen int64
}

type Eighteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	Seventeen
	Eighteen []int
}

type Nineteen struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	Eighteen
	Nineteen []int8
}

type Twenty struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	Nineteen
	Twenty []int16
}

type TwentyOne struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	Twenty
	TwentyOne []int32
}

type TwentyTwo struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	TwentyOne
	TwentyTwo []int64
}

type TwentyThree struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	TwentyTwo
	TwentyThree uint
}

type TwentyFour struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	TwentyThree
	TwentyFour uint16
}

type TwentyFive struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	TwentyFour
	TwentyFive uint32
}

type TwentySix struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	TwentyFive
	TwentySix uint64
}

type TwentySeven struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	TwentySix
	TwentySeven []uint
}

type TwentyEight struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	TwentySeven
	TwentyEight []uint16
}

type TwentyNine struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	TwentyEight
	TwentyNine []uint32
}

type Thirty struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	TwentyNine
	Thirty []uint64
}

type ThirtyOne struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	Thirty
	ThirtyOne string
}

type ThirtyTwo struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	//Thirty
	ThirtyOne
	ThirtyTwo []string
}

type ThirtyThree struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	//Thirty
	//ThirtyOne
	ThirtyTwo
	ThirtyThree time.Time
}

type ThirtyFour struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	//Thirty
	//ThirtyOne
	//ThirtyTwo
	ThirtyThree
	ThirtyFour []time.Time
}

type ThirtyFive struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	//Thirty
	//ThirtyOne
	//ThirtyTwo
	//ThirtyThree
	ThirtyFour
	ThirtyFive time.Duration
}

type ThirtySix struct {
	//One
	//Two
	//Three
	//Four
	//Five
	//Six
	//Seven
	//Eight
	//Nine
	//Ten
	//Eleven
	//Twelve
	//Thirteen
	//Fourteen
	//Fifteen
	//Sixteen
	//Seventeen
	//Eighteen
	//Nineteen
	//Twenty
	//TwentyOne
	//TwentyTwo
	//TwentyThree
	//TwentyFour
	//TwentyFive
	//TwentySix
	//TwentySeven
	//TwentyEight
	//TwentyNine
	//Thirty
	//ThirtyOne
	//ThirtyTwo
	//ThirtyThree
	//ThirtyFour
	ThirtyFive
	ThirtySix []time.Duration
}
