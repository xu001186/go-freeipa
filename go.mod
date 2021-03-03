module github.com/tehwalris/go-freeipa

go 1.15

require (
	github.com/jcmturner/gokrb5/v8 v8.4.2
	github.com/pkg/errors v0.9.1
)

replace github.com/jcmturner/gokrb5 => github.com/ccin2p3/gokrb5 v8.4.3-0.20210303092053-b7136560597c+incompatible
