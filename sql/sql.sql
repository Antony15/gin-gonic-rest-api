-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 18, 2020 at 07:36 PM
-- Server version: 5.7.29-0ubuntu0.18.04.1
-- PHP Version: 5.6.40-8+ubuntu18.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `wednesday-go-test`
--

-- --------------------------------------------------------

--
-- Table structure for table `cabs`
--

CREATE TABLE `cabs` (
  `id` int(11) NOT NULL,
  `cab_id` int(11) NOT NULL,
  `lattitude` float NOT NULL,
  `longitude` float NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `cabs`
--

INSERT INTO `cabs` (`id`, `cab_id`, `lattitude`, `longitude`, `created_at`, `updated_at`) VALUES
(1, 1, 11.0894, 77.0147, '2020-07-18 15:41:40', '2020-07-18 14:15:34'),
(2, 2, 11.0937, 77.0107, '2020-07-18 18:51:38', '2020-07-18 09:51:00');

-- --------------------------------------------------------

--
-- Table structure for table `cab_driver_details`
--

CREATE TABLE `cab_driver_details` (
  `id` int(11) NOT NULL,
  `cab_id` int(11) NOT NULL,
  `driverNname` varchar(50) NOT NULL,
  `vehicleNo` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `cab_driver_details`
--

INSERT INTO `cab_driver_details` (`id`, `cab_id`, `driverNname`, `vehicleNo`, `created_at`) VALUES
(1, 1, 'Kevin', 'TN 37 DA 1234', '2020-07-18 15:43:05'),
(2, 2, 'Naveen', 'TN 38 SK 3489', '2020-07-18 18:49:43');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `address` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `address`) VALUES
(1, 'anto', '123456', 'CBE');

-- --------------------------------------------------------

--
-- Table structure for table `user_bookings`
--

CREATE TABLE `user_bookings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `source` varchar(50) NOT NULL,
  `destination` varchar(50) NOT NULL,
  `total_amount` float NOT NULL,
  `discount_amount` float NOT NULL,
  `actual_amount` float NOT NULL,
  `coupon_applied` int(11) NOT NULL,
  `coupon_id` varchar(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user_bookings`
--

INSERT INTO `user_bookings` (`id`, `user_id`, `source`, `destination`, `total_amount`, `discount_amount`, `actual_amount`, `coupon_applied`, `coupon_id`, `created_at`) VALUES
(1, 1, 'Gandhipuram, Coimbatore', 'Saravanampatty, Coimbatore', 78, 20, 58, 1, 'WEDNESDAY', '2020-07-18 10:01:51'),
(2, 1, 'Peelamedu, Coimbatore', 'Podanur, Coimbatore', 45, 0, 45, 0, '', '2020-07-18 12:02:36');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cabs`
--
ALTER TABLE `cabs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `cab_driver_details`
--
ALTER TABLE `cab_driver_details`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_bookings`
--
ALTER TABLE `user_bookings`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cabs`
--
ALTER TABLE `cabs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `cab_driver_details`
--
ALTER TABLE `cab_driver_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `user_bookings`
--
ALTER TABLE `user_bookings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
