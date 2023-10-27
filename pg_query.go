package pg_query

import (
	"unsafe"

	proto "github.com/golang/protobuf/proto"

	"github.com/pganalyze/pg_query_go/v4/parser"
)

func Scan(input string) (result *ScanResult, err error) {
	protobufScan, err := parser.ScanToProtobuf(input)
	if err != nil {
		return
	}
	result = &ScanResult{}
	err = proto.Unmarshal(protobufScan, result)
	return
}

type SplitResult struct {
	Location int32
	Length   int32
}

func SplitWithScanner(input string) (ret []SplitResult, err error) {
	rst, n, err := parser.SplitWithScanner(input)
	if err != nil {
		return nil, err
	}
	for i := 0; i < n; i++ {
		rptr := (unsafe.Pointer)(uintptr(unsafe.Pointer(&rst[0])) + uintptr(i)*8)
		ptr := (*SplitResult)(rptr)
		ret = append(ret, *ptr)
	}
	return ret, nil
}

// ParseToJSON - Parses the given SQL statement into a parse tree (JSON format)
func ParseToJSON(input string) (result string, err error) {
	return parser.ParseToJSON(input)
}

// Parse the given SQL statement into a parse tree (Go struct format)
func Parse(input string) (tree *ParseResult, err error) {
	protobufTree, err := parser.ParseToProtobuf(input)
	if err != nil {
		return
	}

	tree = &ParseResult{}
	err = proto.Unmarshal(protobufTree, tree)
	return
}

// Parse the given SQL statement into a parse tree (Go struct format)
func ParseWithErrorHandler(input string) (tree *ParseResult, err error) {
	protobufTree, err := parser.ParseToProtobuf(input)
	if err != nil {
		return
	}

	tree = &ParseResult{}
	err = proto.Unmarshal(protobufTree, tree)
	return
}

// Deparses a given Go parse tree into a SQL statement
func Deparse(tree *ParseResult) (output string, err error) {
	protobufTree, err := proto.Marshal(tree)
	if err != nil {
		return
	}

	output, err = parser.DeparseFromProtobuf(protobufTree)
	return
}

// ParsePlPgSqlToJSON - Parses the given PL/pgSQL function statement into a parse tree (JSON format)
func ParsePlPgSqlToJSON(input string) (result string, err error) {
	return parser.ParsePlPgSqlToJSON(input)
}

// Normalize the passed SQL statement to replace constant values with ? characters
func Normalize(input string) (result string, err error) {
	return parser.Normalize(input)
}

// Fingerprint - Fingerprint the passed SQL statement to a hex string
func Fingerprint(input string) (result string, err error) {
	return parser.FingerprintToHexStr(input)
}

// FingerprintToUInt64 - Fingerprint the passed SQL statement to a uint64
func FingerprintToUInt64(input string) (result uint64, err error) {
	return parser.FingerprintToUInt64(input)
}

// HashXXH3_64 - Helper method to run XXH3 hash function (64-bit variant) on the given bytes, with the specified seed
func HashXXH3_64(input []byte, seed uint64) (result uint64) {
	return parser.HashXXH3_64(input, seed)
}
