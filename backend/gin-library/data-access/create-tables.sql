DROP TABLE IF EXISTS gins;
CREATE TABLE gins (
                      id                  INT AUTO_INCREMENT NOT NULL,
                      brandName           VARCHAR(128) NOT NULL,
                      ginName             VARCHAR(255) NOT NULL,
                      ginType             VARCHAR(128) NOT NULL,
                      abv                 FLOAT NOT NULL,                    -- Alcohol By Volume
                      botanicals          JSON,                              -- Key botanicals list stored as JSON
                      distilleryLoc       VARCHAR(255),                      -- Distillery location
                      tastingNotes        TEXT,                              -- Flavor profile description
                      bottleSize          VARCHAR(32),                       -- Bottle size (e.g., 700ml)
                      price               FLOAT,                             -- Price of the gin
                      yearIntroduced      INT,                               -- Year introduced
                      specialEdition      BOOLEAN DEFAULT FALSE,             -- Special/seasonal edition flag
                      awards              JSON,                              -- Awards list stored as JSON
                      distillationMethod  VARCHAR(255),                      -- Distillation method
                      servingSuggestions  TEXT,                              -- Serving suggestions
                      PRIMARY KEY (id)
);