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

### 2014-11-24

BenchmarkJsonMarshalPrimitive	 1000000	      1739 ns/op	     331 B/op	       3 allocs/op
BenchmarkJsonMarshalData	   20000	    100777 ns/op	   71376 B/op	       8 allocs/op
BenchmarkJsonMarshalMap	   10000	    137907 ns/op	   52162 B/op	     312 allocs/op
BenchmarkJsonUnmarshalPrimitive	  500000	      3314 ns/op	     346 B/op	       4 allocs/op
BenchmarkJsonUnmarshalData	    5000	    553074 ns/op	   41523 B/op	       5 allocs/op
BenchmarkJsonUnmarshalMap	    5000	    329064 ns/op	   46649 B/op	     716 allocs/op
BenchmarkGobMarshalPrimitive	  500000	      6595 ns/op	    1517 B/op	      19 allocs/op
BenchmarkGobMarshalData	  200000	     10044 ns/op	   34364 B/op	      16 allocs/op
BenchmarkGobMarshalMap	   10000	    102738 ns/op	   76907 B/op	     524 allocs/op
BenchmarkGobUnmarshalPrimitive	   50000	     51646 ns/op	   17518 B/op	     338 allocs/op
BenchmarkGobUnmarshalData	   50000	     59847 ns/op	   60386 B/op	     325 allocs/op
BenchmarkGobUnmarshalMap	   10000	    166337 ns/op	  101190 B/op	    1146 allocs/op
BenchmarkMsgpMarshalPrimitive	 2000000	       789 ns/op	      96 B/op	       1 allocs/op
BenchmarkMsgpMarshalData	  500000	      5255 ns/op	   16640 B/op	       1 allocs/op
BenchmarkMsgpMarshalMap	  100000	     19564 ns/op	   21248 B/op	       1 allocs/op
BenchmarkMsgpUnmarshalPrimitive	10000000	       253 ns/op	      16 B/op	       1 allocs/op
BenchmarkMsgpUnmarshalData	  500000	      9285 ns/op	   16392 B/op	       1 allocs/op
BenchmarkMsgpUnmarshalMap	   50000	     28376 ns/op	   27881 B/op	     205 allocs/op
BenchmarkProtobufMarshalPrimitive	  200000	      8299 ns/op	     314 B/op	       3 allocs/op
BenchmarkProtobufMarshalData	  100000	     15878 ns/op	   16874 B/op	       3 allocs/op
BenchmarkProtobufMarshalMap	   10000	    108184 ns/op	   95082 B/op	      15 allocs/op
BenchmarkProtobufUnmarshalPrimitive	 2000000	       878 ns/op	     294 B/op	       5 allocs/op
BenchmarkProtobufUnmarshalData	  500000	      6163 ns/op	   16620 B/op	       3 allocs/op
BenchmarkProtobufUnmarshalMap	   50000	     55673 ns/op	   33032 B/op	     510 allocs/op
ok  	github.com/mcellis33/go_serialization_benchmarks	101.796s
