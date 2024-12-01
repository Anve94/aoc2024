module main

replace helper/parser => ../../helper/parser

replace anve/algorithm => ../../algorithm

require (
	anve/algorithm v0.0.0-00010101000000-000000000000
	helper/parser v0.0.0-00010101000000-000000000000
)

go 1.23.3
