package tabungan

import (
	"api-tabungan/domain/tabungan/feature"
)

type TabunganHandler interface {
}

type tabunganHandler struct {
	tabunganFeature feature.TabunganFeature
}

func NewTabunganHandler(tabunganFeature feature.TabunganFeature) TabunganHandler {
	return &tabunganHandler{
		tabunganFeature: tabunganFeature,
	}
}
