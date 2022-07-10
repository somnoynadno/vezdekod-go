package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"unicode"
	"vezdekod-go/50/proto_out"
)

// ReadProtobufFromBinaryFile reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "",
		OrigName:     false,
	}

	return marshaller.MarshalToString(message)
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	cities := &proto_out.Cities{}
	person := &proto_out.Person{}
	points := &proto_out.Points{}
	names := &proto_out.Names{}
	teams := &proto_out.Teams{}

	//filenames := []string{"pb/example1.pb", "pb/example2.pb", "pb/example3.pb", "pb/example4.pb"}
	//structs := []proto.Message{cities, person, points, teams, names}
	//
	//for _, file := range filenames {
	//	for _, s := range structs {
	//		err := ReadProtobufFromBinaryFile(file, s)
	//		if err != nil {
	//			continue
	//		}
	//
	//		j, err := ProtobufToJSON(s)
	//		if err != nil {
	//			continue
	//		}
	//
	//		if isASCII(j) {
	//			fmt.Println(s)
	//		}
	//	}
	//}

	err := ReadProtobufFromBinaryFile("pb/example1.pb", cities)
	if err != nil {
		panic(err)
	}

	fmt.Println(cities)

	err = ReadProtobufFromBinaryFile("pb/example1.pb", names)
	if err != nil {
		panic(err)
	}

	fmt.Println(names)

	err = ReadProtobufFromBinaryFile("pb/example2.pb", points)
	if err != nil {
		panic(err)
	}

	fmt.Println(points)

	err = ReadProtobufFromBinaryFile("pb/example3.pb", person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)

	err = ReadProtobufFromBinaryFile("pb/example4.pb", teams)
	if err != nil {
		panic(err)
	}

	fmt.Println(teams)
}
