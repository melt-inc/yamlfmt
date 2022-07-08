module github.com/jamesrom/yamlfmt

go 1.18

require (
	github.com/spf13/pflag v1.0.5
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/kr/pretty v0.3.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

// Once https://github.com/kubernetes-sigs/yaml/pull/76 lands, remove this
replace sigs.k8s.io/yaml v1.3.0 => github.com/natasha41575/yaml-1 v1.3.1-0.20220514005426-0e00b683066c
