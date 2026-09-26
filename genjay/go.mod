module github.com/speedyhoon/jay/genjay

go 1.27.0

require (
	github.com/dave/dst v0.28.0
	github.com/go-openapi/testify/v2 v2.8.0
	github.com/speedyhoon/ext v0.0.0-20260830071111-64b2ed40796f
	github.com/speedyhoon/flag v0.0.0-20260908115705-80f911a13702
	github.com/speedyhoon/jay v0.0.0
	github.com/speedyhoon/rando v0.0.0-20260926043144-3117ad806214
	github.com/speedyhoon/rando/types v0.0.0-20260926043144-3117ad806214
	github.com/speedyhoon/utl v0.0.0-20260911125320-67808142ca88
	golang.org/x/tools v0.50.0
	mvdan.cc/gofumpt v0.8.0
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/speedyhoon/numnam v0.0.0-20260911154813-7b7cb82ed61c // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
)

replace github.com/speedyhoon/jay => ../
