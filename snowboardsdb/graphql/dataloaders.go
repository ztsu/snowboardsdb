package graphql

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboardsdb/dataloader"
	"net/http"
)

const dataLoadersKey = "dataloaders"

type loaders struct {
	Brands          *dataloader.BrandLoader
	Persons         *dataloader.PersonLoader
	BrandCatalogues *dataloader.BrandCataloguesLoader
}

func ContextWithDataLoaders(services *Stores, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), dataLoadersKey, &loaders{
			Brands:          dataloader.NewBrandLoader(dataloader.NewBrandLoaderConfig(r.Context(), services.Brands)),
			Persons:         dataloader.NewPersonLoader(dataloader.NewPersonLoaderConfig(r.Context(), services.Persons)),
			BrandCatalogues: dataloader.NewBrandCataloguesLoader(dataloader.NewBrandCataloguesLoaderConfig(r.Context(), services.Catalogues)),
		})

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func DataLoaders(ctx context.Context) *loaders {
	return ctx.Value(dataLoadersKey).(*loaders)
}
