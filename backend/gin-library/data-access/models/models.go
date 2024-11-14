package models

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