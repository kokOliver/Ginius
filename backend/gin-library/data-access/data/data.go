package data

import (
	"ginius/gin-library/data-access/models"
)

// gin slice to hold all gin entries
var Gins = []models.GinEntry{
	{
        GinID:            1,
        BrandName:        "Hendrick's",
        GinName:          "Hendrick's Orbium",
        GinType:	      "London Dry",
        ABV:              43.4,
        Botanicals:       []string{"juniper", "rose", "cucumber"},
        DistilleryLoc:    "Scotland",
        TastingNotes:     "Floral, cucumber, and a hint of spice.",
        BottleSize:       "700ml",
        Price:            45.99,
        YearIntroduced:   2018,
        SpecialEdition:   true,
        Awards:           []string{"Gold at IWSC 2019"},
        DistillationMethod: "Pot distillation",
        ServingSuggestions: "Best served with tonic and cucumber garnish.",
    },
	{
		GinID:            2,
		BrandName:        "Tanqueray",
		GinName:          "Tanqueray No. Ten",
		GinType:	      "London Dry",
		ABV:              47.3,
		Botanicals:       []string{"juniper", "coriander", "angelica"},
		DistilleryLoc:    "Scotland",
		TastingNotes:     "Citrus, juniper, and a hint of spice.",
		BottleSize:       "700ml",
		Price:            35.99,
		YearIntroduced:   2000,
		SpecialEdition:   false,
		Awards:           []string{"Gold at IWSC 2018"},
		DistillationMethod: "Pot distillation",
		ServingSuggestions: "Best served with tonic and a slice of grapefruit.",
	},
}