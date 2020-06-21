package main

type Label struct {
	ID      string `json:"id"`
	Aliases []struct {
		Begin    interface{} `json:"begin"`
		End      interface{} `json:"end"`
		Ended    bool        `json:"ended"`
		Locale   string      `json:"locale"`
		Name     string      `json:"name"`
		Primary  bool        `json:"primary"`
		SortName string      `json:"sort-name"`
		Type     string      `json:"type"`
		TypeID   string      `json:"type-id"`
	} `json:"aliases"`
	Country        string        `json:"country"`
	Disambiguation string        `json:"disambiguation"`
	Ipis           []interface{} `json:"ipis"`
	Isnis          []interface{} `json:"isnis"`
	LabelCode      int64         `json:"label-code"`
	Name           string        `json:"name"`
	SortName       string        `json:"sort-name"`
	Type           string        `json:"type"`
	TypeID         string        `json:"type-id"`
}

type Release struct {
	ID           string `json:"id"`
	ArtistCredit []struct {
		Artist struct {
			ID             string `json:"id"`
			Disambiguation string `json:"disambiguation"`
			Name           string `json:"name"`
			SortName       string `json:"sort-name"`
		} `json:"artist"`
		Joinphrase string `json:"joinphrase"`
		Name       string `json:"name"`
	} `json:"artist-credit"`
	Asin            interface{} `json:"asin"`
	Barcode         interface{} `json:"barcode"`
	Country         string      `json:"country"`
	CoverArtArchive struct {
		Artwork  bool  `json:"artwork"`
		Back     bool  `json:"back"`
		Count    int64 `json:"count"`
		Darkened bool  `json:"darkened"`
		Front    bool  `json:"front"`
	} `json:"cover-art-archive"`
	Date           string `json:"date"`
	Disambiguation string `json:"disambiguation"`
	LabelInfo      []struct {
		CatalogNumber string `json:"catalog-number"`
		Label         struct {
			ID             string      `json:"id"`
			Disambiguation string      `json:"disambiguation"`
			LabelCode      interface{} `json:"label-code"`
			Name           string      `json:"name"`
			SortName       string      `json:"sort-name"`
		} `json:"label"`
	} `json:"label-info"`
	Media []struct {
		Discs []struct {
			ID          string  `json:"id"`
			OffsetCount int64   `json:"offset-count"`
			Offsets     []int64 `json:"offsets"`
			Sectors     int64   `json:"sectors"`
		} `json:"discs"`
		Format      string `json:"format"`
		FormatID    string `json:"format-id"`
		Position    int64  `json:"position"`
		Title       string `json:"title"`
		TrackCount  int64  `json:"track-count"`
		TrackOffset int64  `json:"track-offset"`
		Tracks      []struct {
			ID           string `json:"id"`
			ArtistCredit []struct {
				Artist struct {
					ID             string `json:"id"`
					Disambiguation string `json:"disambiguation"`
					Name           string `json:"name"`
					SortName       string `json:"sort-name"`
				} `json:"artist"`
				Joinphrase string `json:"joinphrase"`
				Name       string `json:"name"`
			} `json:"artist-credit"`
			Length    int64 `json:"length"`
			Number    int64 `json:"number,string"`
			Position  int64 `json:"position"`
			Recording struct {
				ID           string `json:"id"`
				ArtistCredit []struct {
					Artist struct {
						ID             string `json:"id"`
						Disambiguation string `json:"disambiguation"`
						Name           string `json:"name"`
						SortName       string `json:"sort-name"`
					} `json:"artist"`
					Joinphrase string `json:"joinphrase"`
					Name       string `json:"name"`
				} `json:"artist-credit"`
				Disambiguation string `json:"disambiguation"`
				Length         int64  `json:"length"`
				Title          string `json:"title"`
				Video          bool   `json:"video"`
			} `json:"recording"`
			Title string `json:"title"`
		} `json:"tracks"`
	} `json:"media"`
	Packaging     interface{} `json:"packaging"`
	PackagingID   interface{} `json:"packaging-id"`
	Quality       string      `json:"quality"`
	ReleaseEvents []struct {
		Area struct {
			ID             string   `json:"id"`
			Disambiguation string   `json:"disambiguation"`
			Iso31661Codes  []string `json:"iso-3166-1-codes"`
			Name           string   `json:"name"`
			SortName       string   `json:"sort-name"`
		} `json:"area"`
		Date string `json:"date"`
	} `json:"release-events"`
	Status             string `json:"status"`
	StatusID           string `json:"status-id"`
	TextRepresentation struct {
		Language string `json:"language"`
		Script   string `json:"script"`
	} `json:"text-representation"`
	Title string `json:"title"`
}
