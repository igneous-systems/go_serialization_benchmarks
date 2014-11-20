## Reading Results

This file contains the direct output from `go test`. This includes the following columns, in order:

1. benchmark name
2. number of iterations
3. average time per iteration
4. average bytes allocated per iteration
5. average number of allocations per iteration

## Notes

- JSON serialization of byte slices is very slow, because everything, including byte slices, needs to be encoded as a string. The other formats store byte slices raw.
- Gob results are not really representative of normal performance. Gob is designed for serializing streams or vectors of a single type, not individual values. There is some expensive reflection that takes place the first time a particular type is serialized/deserialized that does not take place for subsequent serializations/deserializations of the same type.
- msgp appears to be the most efficient of the tested options.

## Results

### 2014-11-19

BenchmarkJsonMarshalPrimitive	  500000	      4564 ns/op	     756 B/op	       8 allocs/op
BenchmarkJsonMarshalData	   10000	    114656 ns/op	   71743 B/op	      12 allocs/op
BenchmarkJsonMarshalMap	    2000	    809388 ns/op	  308946 B/op	     321 allocs/op
BenchmarkJsonUnmarshalPrimitive	  500000	      7069 ns/op	     654 B/op	      12 allocs/op
BenchmarkJsonUnmarshalData	    5000	    547484 ns/op	   41771 B/op	      12 allocs/op
BenchmarkJsonUnmarshalMap	    1000	   2174408 ns/op	  191261 B/op	    3713 allocs/op
BenchmarkGobMarshalPrimitive	  200000	      9755 ns/op	    1684 B/op	      23 allocs/op
BenchmarkGobMarshalData	  200000	     13256 ns/op	   34537 B/op	      20 allocs/op
BenchmarkGobMarshalMap	   10000	    128557 ns/op	   77092 B/op	     528 allocs/op
BenchmarkGobUnmarshalPrimitive	   50000	     69697 ns/op	   17020 B/op	     342 allocs/op
BenchmarkGobUnmarshalData	   50000	     56522 ns/op	   44251 B/op	     329 allocs/op
BenchmarkGobUnmarshalMap	   10000	    177545 ns/op	   93017 B/op	    1141 allocs/op
BenchmarkMsgpMarshalPrimitive	 1000000	      2568 ns/op	     265 B/op	       5 allocs/op
BenchmarkMsgpMarshalData	  500000	      5900 ns/op	   16809 B/op	       5 allocs/op
BenchmarkMsgpMarshalMap	  100000	     24728 ns/op	   21417 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalPrimitive	 1000000	      2437 ns/op	     186 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalData	  500000	      3185 ns/op	     178 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalMap	   50000	     42409 ns/op	   22570 B/op	     204 allocs/op
BenchmarkProtobufMarshalPrimitive	  500000	      3313 ns/op	     485 B/op	       8 allocs/op
BenchmarkProtobufMarshalData	  500000	      6791 ns/op	   17044 B/op	       8 allocs/op
BenchmarkProtobufMarshalMap	   20000	     80245 ns/op	   95251 B/op	      20 allocs/op
BenchmarkProtobufUnmarshalPrimitive	  500000	      3188 ns/op	     465 B/op	       9 allocs/op
BenchmarkProtobufUnmarshalData	  500000	      5968 ns/op	   16790 B/op	       8 allocs/op
BenchmarkProtobufUnmarshalMap	   50000	     60525 ns/op	   33211 B/op	     514 allocs/op
ok  	github.com/mcellis33/go_serialization_benchmarks	140.236s
