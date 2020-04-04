package custom

// model of manifest file actually reason for doing that package
type model struct {
	Frames map[string] /*imgName*/ frameInstance `json:"frames"`
	Meta   meta                                  `json:"meta"`
}

type frameInstance struct {
	// position of target image on original sprite shit
	Frame frame `json:"frame"`

	Rotated bool `json:"rotated"`
	Trimmed bool `json:"trimmed"`

	// original Sprite rectangle values
	SpriteSourceSize frame `json:"spriteSourceSize"`

	// original file weight / height proportions
	SourceSize wh `json:"source_size"`
}

type frame struct {
	X int `json:"x"`
	Y int `json:"y"`
	wh
}

type wh struct {
	W int `json:"w"`
	H int `json:"h"`
}

type meta struct {
	Version string `json:"version"`
	Image   string `json:"image"`
	Format  string `json:"format"`
	Size    wh     `json:"size"`
	Scale   string `json:"scale"`
}
