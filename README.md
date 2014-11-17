# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested Serialization Methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/philhofer/msgp](https://github.com/philhofer/msgp) *(code generator for msgpack)*
- [code.google.com/p/goprotobuf/proto](https://code.google.com/p/protobuf)

## TODO Serialization Methods

- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/pquerna/ffjson](https://github.com/pquerna/ffjson)

## Untested Serialization Methods

- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*. This package is too large to download in a reasonable amount of time. Given the package layout problems, I assume the package is a mess.
- [github.com/ugorji/go-msgpack](https://github.com/ugorji/go-msgpack) deprecated

## Updating Structs

Both msgp and protobuf require code generation from the structs that will be serialized. msgp generates methods for the existing struct from a given file. protobuf generates structs from the .proto file. Whenever you change structdef.go, you will need to make the equivalent change in structdef.proto and run the preprocessors.

- msgp: `msgp -file=structdef.go -marshal=true -o=structdef.mp.go` generates structdef.mp.go and structdef.mp_test.go. You can delete structdef.mp_test.go, as it simply contains superfluous unit tests and benchmarks for the generated methods.
- protobuf: `protoc --go_out=. structdef.proto` generates structdef.pb.go

## Running the benchmarks

```bash
go get -u -t
go test -bench='.*' ./
```
Publish interesting results to results.md in this repo.

## Validate

The benchmarks can also be run with validation enabled. This compares the input to Marshal with the output from Unmarshal in each unmarshaling test.

```bash
VALIDATE=1 go test -bench='.*' ./
```

## Data

Serialization for the following canonical data structures is tested:

1. Primitive: this struct contains only primitive types.
2. Data: this struct contains a short metadata string as well as a 16KiB []byte data segment.
3. Map: this struct contains a map from string to string with 100 entries with key and value size 100.

## Time Issues

There used to be a test that serialized the time.Time type in this package. The following issues with serializing time.Time were found using VALIDATE.

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{3b86c4a97a5aa287 2014-09-26 14:46:15.684430354 +1000 AEST a3ff184699 4 true 0.5503346859316104}
        &{3b86c4a97a5aa287 2014-09-26 14:46:15.684 +1000 AEST a3ff184699 4 true 0.5503346859316104}
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{825f2ed8bc78185b 2014-09-26 14:46:17.126931876 +1000 AEST 929f58adf2 4 true 0.19285299476299536}
        &{825f2ed8bc78185b 2014-09-26 04:46:17.126 +0000 UTC 929f58adf2 4 true 0.19285299476299536}
BenchmarkUgorjiCodecBincUnmarshal   --- FAIL: BenchmarkUgorjiCodecBincUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{8ca5570b13d51126 2014-09-26 14:46:35.800449873 +1000 AEST 89522df312 2 false 0.6136756208926619}
        &{8ca5570b13d51126 2014-09-26 14:46:35.800449873 +1000 +1000 89522df312 2 false 0.6136756208926619}
```
