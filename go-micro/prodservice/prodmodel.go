package prodservice

import "strconv"

type ProdModel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProdModel {
	result := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		result = append(result, NewProd(100+i, "prodname"+strconv.Itoa(100+i)))
	}
	return result
}
