CREATE TABLE IF NOT EXISTS `genres` (
  `genreID` varchar(256),
  `displayName` TEXT,
  PRIMARY KEY (`genreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `subGenres` (
  `subGenreID` varchar(256),
  `displayName` varchar(256),
  PRIMARY KEY (`subGenreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `series` (
  `seriesID` varchar(256),
  `displayName` varchar(256),
  `description` TEXT,
  `imageURL` TEXT,
  `genreID` varchar(256),
  PRIMARY KEY (`seriesID`),
  CONSTRAINT `fk_series_genre` FOREIGN KEY (`genreID`) REFERENCES `genres` (`genreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `seriesSubGenres` (
  `seriesID` varchar(256),
  `subGenreID` varchar(256),
  PRIMARY KEY (`seriesID`, `subGenreID`),
  CONSTRAINT `rel_series_sub_genres_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`),
  CONSTRAINT `rel_series_sub_genres_sub_genre` FOREIGN KEY (`subGenreID`) REFERENCES `subGenres` (`subGenreID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `seasons` (
  `seasonID` varchar(256),
  `seriesID` varchar(256) NOT NULL,
  `displayName` varchar(256),
  `imageURL` TEXT,
  `displayOrder` int,
  PRIMARY KEY (`seasonID`),
  CONSTRAINT `fk_season_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `episodes` (
  `episodeID` varchar(256),
  `seasonID` varchar(256),
  `seriesID` varchar(256) NOT NULL,
  `displayName` varchar(256),
  `description` varchar(256),
  `imageURL` varchar(256),
  `displayOrder` int,
  PRIMARY KEY (`episodeID`),
  CONSTRAINT `fk_episode_seasons` FOREIGN KEY (`seasonID`) REFERENCES `seasons` (`seasonID`) ON DELETE SET NULL,
  CONSTRAINT `fk_episode_series` FOREIGN KEY (`seriesID`) REFERENCES `series` (`seriesID`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
