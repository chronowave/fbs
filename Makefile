generate: generate_go generate_rust

generate_go:
	cd go && flatc --go --go-namespace fbs ../lib.fbs && mv fbs/*.go ./ && rm -r fbs && go build

generate_rust:
	cd rust && cargo build

clean:
	rm -rf go/*.go
	rm -rf rust/src/lib.rs
	touch rust/src/lib.rs
