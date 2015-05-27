all: read interface counter noticias

read: read/read.rb read/read.py read/read.go read/Read.java read/read.c

interface: interfaces/inter.go

counter: counter/counter1.go counter/counter2.go counter/counter3.go

noticias: noticias/noticias1.go

noticias_server:
	go run noticias/noticias3.go

*/*.rb: .update
	time ruby $@

*/*.py: .update
	time python $@

*/*.go: .update
	@go build -o bin/$(basename $@)_go $@
	time ./bin/$(basename $@)_go

*/*.java: .update
	@javac $@
	time cd read && java $(basename $(notdir $@))

*/*.c: .update
	@cc -o bin/$(basename $@)_c -O3 $@
	time ./bin/$(basename $@)_c

.update:

.PHONY: .update all read interface counter noticias
