module github.com/speedyhoon/jay/genjay

go 1.24.5

require (
	github.com/dave/dst v0.27.3
	github.com/go-openapi/testify/v2 v2.4.1
	github.com/speedyhoon/ext v0.0.0-20260113124402-f0066d60c0b7
	github.com/speedyhoon/flag v0.0.0-20260227112925-86bbdde51b02
	github.com/speedyhoon/jay v0.0.0-00010101000000-000000000000
	github.com/speedyhoon/rando v0.0.0-20260324092358-d3a30fdead04
	github.com/speedyhoon/rando/types v0.0.0-20260324092358-d3a30fdead04
	github.com/speedyhoon/utl v0.0.0-20260315024855-0c863838fe45
	golang.org/x/tools v0.42.0
	mvdan.cc/gofumpt v0.8.0
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/speedyhoon/numnam v0.0.0-20260203072053-447015b4d8d5 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
)

replace github.com/speedyhoon/jay => ../
