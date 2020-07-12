package snowboards

type Brand struct {
	ID          int
	Name        string
	WebsiteURL  string
	FoundedIn   *int
	Founders    []int
	OriginsFrom *string
}

type Person struct {
	ID   int
	Name string
}

type Catalogue struct {
	ID      int
	BrandID int
	Season  string
	Type    string // @todo use CatalogueType
	URL     string
	Size    int
}

type CatalogueType string

const (
	CatalogueTypeIssuu CatalogueType = "issuu"
)

type Snowboard struct {
	ID      int
	BrandID int
	Name    string
	Season  string
	Type    SnowboardType
}

type SnowboardType string

const (
	SnowboardTypeSnowboard   = "snowboard"
	SnowboardTypeSplitboard  = "splitboard"
	SnowboardTypePowsurfer   = "powsurfer"
	SnowboardTypeSplitsurfer = "splitsurfer"
	SnowboardTypeSnowskate   = "snowskate"
)

type Image struct {
	ID          int
	SnowboardID int
	URL         string
	Size        *string
	ColorOfBase *string
}
