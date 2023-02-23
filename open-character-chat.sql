-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Feb 02, 2023 at 10:37 AM
-- Server version: 5.7.34
-- PHP Version: 8.0.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `open-character-chat`
--

-- --------------------------------------------------------

--
-- Table structure for table `Conversations`
--

CREATE TABLE `Conversations` (
  `id` int(11) NOT NULL,
  `userId` int(11) NOT NULL,
  `characterId` int(11) NOT NULL,
  `isOpen` tinyint(1) NOT NULL,
  `lastMsgDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Game`
--

CREATE TABLE `Game` (
  `id` int(11) NOT NULL,
  `title` text NOT NULL,
  `releaseDate` text NOT NULL,
  `isMultiplayer` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Game`
--

INSERT INTO `Game` (`id`, `title`, `releaseDate`, `isMultiplayer`) VALUES
(1, 'League of Legends', '01-01-2009', 1),
(2, 'Call of Duty Modern Warfare 2', '01-10-2022', 1),
(3, 'League of Legends', '27-10-2009', 1),
(4, 'League of Legends', '27-10-2009', 0),
(5, 'League of Legends', '27-10-2009', 1);

-- --------------------------------------------------------

--
-- Table structure for table `gameCharacter`
--

CREATE TABLE `gameCharacter` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `personality` text NOT NULL,
  `game` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `gameCharacter`
--

INSERT INTO `gameCharacter` (`id`, `name`, `personality`, `game`) VALUES
(6, 'Gangplank', 'Gangplank was once a feared pirate captain, sailing the seas in search of plunder and profit. His crew was loyal to him, admiring his ruthless cunning and brutality. He eventually became the commander of the 5th Fleet of the Bilgewater Armada and declared himself \'The Saltwater Scourge\'. He was eventually betrayed and killed by his first mate, but his spirit lives on as a champion in League of Legends.', 1);

-- --------------------------------------------------------

--
-- Table structure for table `Message`
--

CREATE TABLE `Message` (
  `id` int(11) NOT NULL,
  `content` text NOT NULL,
  `datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `User`
--

CREATE TABLE `User` (
  `id` int(11) NOT NULL,
  `username` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Conversations`
--
ALTER TABLE `Conversations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `userId` (`userId`),
  ADD KEY `characterId` (`characterId`);

--
-- Indexes for table `Game`
--
ALTER TABLE `Game`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `gameCharacter`
--
ALTER TABLE `gameCharacter`
  ADD PRIMARY KEY (`id`),
  ADD KEY `game` (`game`);

--
-- Indexes for table `Message`
--
ALTER TABLE `Message`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `User`
--
ALTER TABLE `User`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Conversations`
--
ALTER TABLE `Conversations`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Game`
--
ALTER TABLE `Game`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `gameCharacter`
--
ALTER TABLE `gameCharacter`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `Message`
--
ALTER TABLE `Message`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `User`
--
ALTER TABLE `User`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Conversations`
--
ALTER TABLE `Conversations`
  ADD CONSTRAINT `conversations_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `User` (`id`),
  ADD CONSTRAINT `conversations_ibfk_2` FOREIGN KEY (`characterId`) REFERENCES `gameCharacter` (`id`);

--
-- Constraints for table `gameCharacter`
--
ALTER TABLE `gameCharacter`
  ADD CONSTRAINT `gamecharacter_ibfk_1` FOREIGN KEY (`game`) REFERENCES `Game` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
