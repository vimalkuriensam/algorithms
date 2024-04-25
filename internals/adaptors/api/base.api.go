package api

import "github.com/vimalkuriensam/algorithms/internals/ports"

type Adaptor struct {
	sort ports.SortingCorePort
}

func Initialize(sort ports.SortingCorePort) *Adaptor {
	return &Adaptor{
		sort: sort,
	}
}
