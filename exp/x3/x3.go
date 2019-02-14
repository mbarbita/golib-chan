package main

import "fmt"

//controlers map
// var controllers map[int]int
// var controller map[int]int

//[loop id]  = component list id
// var loopMap map[int]int
var inputMap map[int]chan interface{}
var outputMap map[int]chan interface{}

// component type
// 0=dummy; 1=frame; 2=echo; 3=router
// var compType map[int]int

// [comp list id][comp id] = comp type
var compList map[int]map[int]int

func main() {
	compList = make(map[int]map[int]int)
	compList[0] = make(map[int]int)
	compList[0][0] = 2
	compList[0][1] = 3
	compList[0][2] = 3

	compList[1] = make(map[int]int)
	compList[1][0] = 0
	compList[1][1] = 2
	compList[1][2] = 1

	fmt.Printf("compList: %+v\n", compList)

	// loopMap = make(map[int]int)
	// loopMap[0]=0
	// compType = make(map[int]int)
	// compType[0] = 0
	// compType[1] = 1
	// compType[2] = 2
	// compType[3] = 3
	// fmt.Printf("compType: %+v\n", compType)

	// controller1[0] = 0
	// fmt.Printf("controller1: %+v\n", controller1)
	//
	// loopMap[0] = make(map[int]int)
	// loopMap[0][0] = 0
	// fmt.Printf("loopMap: %+v\n", loopMap)
	//
	// fmt.Printf("inputMap: %+v\n", inputMap)
	// fmt.Printf("outputMap: %+v\n", outputMap)
	//
	// // [comp list id][component id] = type
	// compList[0] = make(map[int]int)
	// compList[0][0] = 1
}
