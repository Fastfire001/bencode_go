package bencode

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type bstring struct {
	ItemType reflect.Kind
	Item interface{}
	Bstring String
}

type Int int
type String string

func (i bstring) Show() {
	fmt.Println("{ ItemType:", i.ItemType, ", Item:", i.Item, ", Bstring:", i.Bstring, "}")
}

func Encode (item interface{}) (toReturn bstring) {
	reflectType := reflect.TypeOf(item).Kind()
	encodedItem := String("")
	switch reflectType {
	case reflect.String:
		encodedItem = String(item.(string)).BencodeString()
	case reflect.Int:
		encodedItem = Int(item.(int)).BencodeInt()
	default:
		panic(errors.New("Encode accept int and string"))
	}
	toReturn.Item = item
	toReturn.ItemType = reflectType
	toReturn.Bstring = encodedItem
	return
}

func (i Int) BencodeInt() (bstring String) {
	bstring = String("i" + strconv.Itoa(int(i)) + "e")
	return
}

func (s String) BencodeString() (bstring String) {
	bstring = String(strconv.Itoa(len(s)) + ":") + s
	return
}
