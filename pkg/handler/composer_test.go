package handler_test

import (
	"github.com/RobertVillalba/tusd/pkg/filestore"
	"github.com/RobertVillalba/tusd/pkg/handler"
	"github.com/RobertVillalba/tusd/pkg/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := handler.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	config := handler.Config{
		StoreComposer: composer,
	}

	_, _ = handler.NewHandler(config)
}
