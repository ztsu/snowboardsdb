package graphql

import (
	"context"
	"fmt"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
	"log"
)

//
type snowboardResolver struct {
	*rootResolver
}

func (r *snowboardResolver) FullName(ctx context.Context, obj *Snowboard) (string, error) {
	if b, ok := obj.Brand.(*Brand); !ok {
		return "", nil
	} else {
		brand, err := DataLoaders(ctx).Brands.Load(b.ID)
		if err != nil {
			log.Printf("can't resolve brand: %s", err)
			return "", nil
		}

		if brand == nil {
			return "", nil
		}

		return fmt.Sprintf("%s %s", brand.Name, obj.Name), nil
	}
}

func (r *snowboardResolver) Brand(ctx context.Context, obj *Snowboard) (BrandResolveResult, error) {
	if b, ok := obj.Brand.(*Brand); ok {
		brand, err := DataLoaders(ctx).Brands.Load(b.ID)
		if err != nil {
			log.Printf("can't resolve brand: %s", err)
			return BrandResolveError{Message: "can't resolve brand"}, nil
		}

		if brand == nil {
			return BrandNotFoundError{Message: "brand not found"}, nil
		}

		return brandToGraphQL(brand), nil
	}

	log.Printf("can't resolve brand: obj is not *Brand")

	return BrandResolveError{Message: "can't resolve brand"}, nil
}

func (r *snowboardResolver) Images(ctx context.Context, obj *Snowboard, limit int, offset int) ([]SnowboardImage, error) {
	var (
		limitUint64  = uint64(limit)
		offsetUint64 = uint64(offset)
	)

	query := snowboardsdb.ImageQuery{
		SnowboardID: []int{obj.ID},
		Limit:       &limitUint64,
		Offset:      &offsetUint64,
	}

	images, err := r.Stores.Images.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't resolve images on a snowboard: %w", err)
	}

	items := make([]SnowboardImage, len(images))

	for j, img := range images {
		cg := imageToGraphQL(img)

		items[j] = cg
	}

	return items, nil
}

func imageToGraphQL(i *snowboardsdb.Image) SnowboardImage {
	switch image := i; {
	case image.ColorOfBase != nil:
		return SnowboardBaseImage{
			URL:         i.URL,
			ColorOfBase: *i.ColorOfBase, // @todo: to enum
		}
	case image.Size != nil:
		return SnowboardSizeImage{
			URL:  i.URL,
			Size: *i.Size, // @todo: to enum
		}
	default:
		return SnowboardGeneralImage{
			URL: i.URL,
		}
	}
}
