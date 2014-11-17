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

### 2014-11-18

BenchmarkJsonMarshalPrimitive	  500000	      4488 ns/op	     756 B/op	       8 allocs/op
BenchmarkJsonMarshalData	   20000	     97373 ns/op	   71737 B/op	      12 allocs/op
BenchmarkJsonMarshalMap	    2000	    773419 ns/op	  308684 B/op	     321 allocs/op
BenchmarkJsonUnmarshalPrimitive	  500000	      6829 ns/op	     654 B/op	      12 allocs/op
BenchmarkJsonUnmarshalData	    5000	    546035 ns/op	   41772 B/op	      12 allocs/op
BenchmarkJsonUnmarshalMap	    1000	   2099545 ns/op	  191258 B/op	    3713 allocs/op
BenchmarkGobMarshalPrimitive	  200000	      9492 ns/op	    1685 B/op	      23 allocs/op
BenchmarkGobMarshalData	  200000	     12879 ns/op	   34537 B/op	      20 allocs/op
BenchmarkGobMarshalMap	   10000	    116104 ns/op	   77091 B/op	     528 allocs/op
BenchmarkGobUnmarshalPrimitive	   50000	     67174 ns/op	   17018 B/op	     342 allocs/op
BenchmarkGobUnmarshalData	   50000	     51052 ns/op	   44250 B/op	     329 allocs/op
BenchmarkGobUnmarshalMap	   10000	    170612 ns/op	   93020 B/op	    1141 allocs/op
BenchmarkMsgpMarshalPrimitive	 1000000	      2642 ns/op	     265 B/op	       5 allocs/op
BenchmarkMsgpMarshalData	  500000	      6092 ns/op	   16809 B/op	       5 allocs/op
BenchmarkMsgpMarshalMap	   50000	     25640 ns/op	   21417 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalPrimitive	 1000000	      2544 ns/op	     186 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalData	 1000000	      2953 ns/op	     178 B/op	       5 allocs/op
BenchmarkMsgpUnmarshalMap	   50000	     39245 ns/op	   22570 B/op	     204 allocs/op
BenchmarkProtobufMarshalPrimitive	  500000	      3320 ns/op	     485 B/op	       8 allocs/op
BenchmarkProtobufMarshalData	  500000	      6810 ns/op	   17044 B/op	       8 allocs/op
BenchmarkProtobufMarshalMap	   50000	     76571 ns/op	   95268 B/op	      21 allocs/op
BenchmarkProtobufUnmarshalPrimitive	  500000	      3165 ns/op	     465 B/op	       9 allocs/op
BenchmarkProtobufUnmarshalData	  500000	      5770 ns/op	   16790 B/op	       8 allocs/op
BenchmarkProtobufUnmarshalMap	   50000	     57970 ns/op	   33261 B/op	     515 allocs/op
ok  	github.com/mcellis33/go_serialization_benchmarks	142.733s