DROP TABLE IF EXISTS `episodes`;
DROP TABLE IF EXISTS `seasons`;
DROP TABLE IF EXISTS `seriesSubGenres`;
DROP TABLE IF EXISTS `series`;
DROP TABLE IF EXISTS `subGenres`;
DROP TABLE IF EXISTS `genres`;

CREATE TABLE `genres` (
  `genreID` varchar(256),
  `displayName` TEXT,
  PRIMARY KEY (`genreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `subGenres` (
  `subGenreID` varchar(256),
  `displayName` varchar(256),
  PRIMARY KEY (`subGenreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `series` (
  `seriesID` varchar(256),
  `displayName` varchar(256),
  `description` TEXT,
  `imageURL` TEXT,
  PRIMARY KEY (`seriesID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `seriesSubGenres` (
  `seriesID` varchar(256),
  `subGenreID` varchar(256),
  PRIMARY KEY (`seriesID`, `subGenreID`),
  CONSTRAINT `rel_series_sub_genres_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `rel_series_sub_genres_sub_genre` FOREIGN KEY (`subGenreID`) REFERENCES `subGenres` (`subGenreID`) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `seasons` (
  `seasonID` varchar(256),
  `seriesID` varchar(256) NOT NULL,
  `displayName` varchar(256),
  `imageURL` TEXT,
  `displayOrder` int(10),
  PRIMARY KEY (`seasonID`),
  CONSTRAINT `fk_season_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `episodes` (
  `episodeID` varchar(256),
  `seasonID` varchar(256) DEFAULT NULL,
  `seriesID` varchar(256) NOT NULL,
  `displayName` varchar(256),
  `description` varchar(256),
  `imageURL` varchar(256),
  `displayOrder` int(10),
  PRIMARY KEY (`episodeID`),
  CONSTRAINT `fk_episode_seasons` FOREIGN KEY (`seasonID`) REFERENCES `seasons` (`seasonID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_episode_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
