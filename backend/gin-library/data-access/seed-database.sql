-- noinspection SqlNoDataSourceInspection
INSERT INTO gins (
    id, brandName, ginName, ginType, abv, botanicals, distilleryLoc, tastingNotes, bottleSize, price, yearIntroduced, specialEdition, awards, distillationMethod, servingSuggestions
) VALUES
      (
          1,
          "Hendrick's",
          "Hendrick's Orbium",
          "London Dry",
          43.4,
          JSON_ARRAY("juniper", "rose", "cucumber"),
          "Scotland",
          "Floral, cucumber, and a hint of spice.",
          "700ml",
          45.99,
          2018,
          TRUE,
          JSON_ARRAY("Gold at IWSC 2019"),
          "Pot distillation",
          "Best served with tonic and cucumber garnish."
      ),
      (
          2,
          "Tanqueray",
          "Tanqueray No. Ten",
          "London Dry",
          47.3,
          JSON_ARRAY("juniper", "coriander", "angelica"),
          "Scotland",
          "Citrus, juniper, and a hint of spice.",
          "700ml",
          35.99,
          2000,
          FALSE,
          JSON_ARRAY("Gold at IWSC 2018"),
          "Pot distillation",
          "Best served with tonic and a slice of grapefruit."
      );
