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

### 2014-11-20

BenchmarkJsonMarshalPrimitive	 1000000	      2461 ns/op	     587 B/op	       3 allocs/op
BenchmarkJsonMarshalData	   20000	     99711 ns/op	   71606 B/op	       8 allocs/op
BenchmarkJsonMarshalMap	    2000	    758646 ns/op	  308784 B/op	     316 allocs/op
BenchmarkJsonUnmarshalPrimitive	  500000	      4828 ns/op	     480 B/op	       7 allocs/op
BenchmarkJsonUnmarshalData	    5000	    574865 ns/op	   41600 B/op	       7 allocs/op
BenchmarkJsonUnmarshalMap	    1000	   2112694 ns/op	  191160 B/op	    3710 allocs/op
BenchmarkGobMarshalPrimitive	  500000	      7004 ns/op	    1517 B/op	      19 allocs/op
BenchmarkGobMarshalData	  200000	     10116 ns/op	   34364 B/op	      16 allocs/op
BenchmarkGobMarshalMap	   10000	    116106 ns/op	   76919 B/op	     524 allocs/op
BenchmarkGobUnmarshalPrimitive	   50000	     61256 ns/op	   17352 B/op	     338 allocs/op
BenchmarkGobUnmarshalData	   50000	     49137 ns/op	   44062 B/op	     324 allocs/op
BenchmarkGobUnmarshalMap	   10000	    154952 ns/op	   92847 B/op	    1137 allocs/op
BenchmarkMsgpMarshalPrimitive	 5000000	       483 ns/op	      96 B/op	       1 allocs/op
BenchmarkMsgpMarshalData	  500000	      3858 ns/op	   16640 B/op	       1 allocs/op
BenchmarkMsgpMarshalMap	  100000	     22018 ns/op	   21248 B/op	       1 allocs/op
BenchmarkMsgpUnmarshalPrimitive	 5000000	       327 ns/op	      16 B/op	       1 allocs/op
BenchmarkMsgpUnmarshalData	 1000000	      1117 ns/op	       8 B/op	       0 allocs/op
BenchmarkMsgpUnmarshalMap	   50000	     35383 ns/op	   22400 B/op	     200 allocs/op
BenchmarkProtobufMarshalPrimitive	 1000000	      1234 ns/op	     314 B/op	       3 allocs/op
BenchmarkProtobufMarshalData	  500000	      4723 ns/op	   16874 B/op	       3 allocs/op
BenchmarkProtobufMarshalMap	   20000	     77031 ns/op	   95082 B/op	      15 allocs/op
BenchmarkProtobufUnmarshalPrimitive	 1000000	      1138 ns/op	     293 B/op	       5 allocs/op
BenchmarkProtobufUnmarshalData	  500000	      3995 ns/op	   16619 B/op	       3 allocs/op
BenchmarkProtobufUnmarshalMap	   50000	     53091 ns/op	   33038 B/op	     510 allocs/op
ok  	github.com/mcellis33/go_serialization_benchmarks	94.876s
