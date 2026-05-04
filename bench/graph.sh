go test ./... -bench=. -benchmem > bench.out
cat bench.out | vizb -n "Benchmark comparisons" -l -p x_n_y -c bar,pie -o graph.html -d "Serialisation performance for 24 struct fields of various types using different technologies."
