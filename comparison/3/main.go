package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Document struct {
	Id int
	moreData
}

type moreData struct {
	IsPublic  bool
	Viewcount int
}

func main() {
	extraData := moreData{
		IsPublic:  true,
		Viewcount: 100,
	}
	myDoc := Document{
		Id:       1,
		moreData: extraData,
	}

	myDocBs, _ := json.Marshal(myDoc)

	var myDocUnmarchalled Document
	_ = json.Unmarshal(myDocBs, &myDocUnmarchalled)
	docsAreEqual := reflect.DeepEqual(myDoc, myDocUnmarchalled)
	if !docsAreEqual {
		fmt.Printf("Not equal")
	} else {
		fmt.Printf("Equal\nmyDoc is %v\nmyDocUnmarchalled is %v",
			myDoc, myDocUnmarchalled)
	}
}
