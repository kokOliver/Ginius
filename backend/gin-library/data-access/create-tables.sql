DROP TABLE IF EXISTS gins;
CREATE TABLE gins (
                       id          INT AUTO_INCREMENT NOT NULL,
                       brandName   VARCHAR(128) NOT NULL,
                       ginName     VARCHAR(255) NOT NULL,
                       ginType     VARCHER(128) NOT NULL,
                       PRIMARY KEY (`id`)
);