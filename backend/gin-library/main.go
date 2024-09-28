package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Gin represents the structure of a gin entry in the database
type GinEntry struct {
    GinID          int      `json:"gin_id"`          // Unique identifier
    BrandName      string   `json:"brand_name"`      // Brand name of the gin
    GinName        string   `json:"gin_name"`        // Name of the gin/expression
    GinType        string   `json:"gin_type"`        // Type (e.g., London Dry, Plymouth)
    ABV            float32  `json:"abv"`             // Alcohol By Volume (e.g., 40.5)
    Botanicals     []string `json:"botanicals"`      // List of key botanicals
    DistilleryLoc  string   `json:"distillery_loc"`  // Distillery location
    TastingNotes   string   `json:"tasting_notes"`   // Flavor profile description
    BottleSize     string   `json:"bottle_size"`     // Bottle size (e.g., 700ml)
    Price          float32  `json:"price"`           // Price of the gin
    YearIntroduced int      `json:"year_introduced"` // Year the gin was first introduced
    SpecialEdition bool     `json:"special_edition"` // Indicates if it's a special/seasonal edition
    Awards         []string `json:"awards"`          // List of awards the gin has won
    DistillationMethod string `json:"distillation_method"` // How the gin is distilled
    ServingSuggestions string `json:"serving_suggestions"` // How to serve the gin
}

// gin slice to hold all gin entries
var gins = []GinEntry{
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

func main() {
	router:= gin.Default()
	router.GET("/gins", getGins)
	router.GET("/gins/:id", getGinByID)
	router.POST("/gins", postGin)

	router.Run("localhost:8080")
}

// getGins returns all gin entries
func getGins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gins)
}

// postGin adds a new gin entry from the JSON received in the request body
func postGin(c *gin.Context) {
	var newGin GinEntry

	if err := c.BindJSON(&newGin); err != nil {
		return
	}

	gins = append(gins, newGin)
	c.IndentedJSON(http.StatusCreated, newGin)
}

// getGinByID returns a gin entry by its ID
func getGinByID(c *gin.Context) {
	id := c.Param("id")
	ginID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid gin ID"})
		return
	}

	for _, gin := range gins {
		if gin.GinID == ginID {
			c.IndentedJSON(http.StatusOK, gin)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "gin not found"})
}
