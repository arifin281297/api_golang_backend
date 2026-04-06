-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 27 Mar 2026 pada 11.06
-- Versi server: 10.4.17-MariaDB
-- Versi PHP: 7.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crudapigolang`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `t_user`
--

DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `cabang_id` bigint(20) DEFAULT 0,
  `group_name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `t_user`
--

INSERT INTO `t_user` (`id`, `username`, `password`, `name`, `email`, `cabang_id`, `group_name`) VALUES
(6, 'esti', '1234', 'Esti', 'esti@gmail.com', 1, 'Admin'),
(7, 'zainul', '1234', 'Ahmad Zainul Arifin', 'zainul@example.com', 2, 'Admin'),
(8, 'keifano', '1234', 'kei', 'keifano@example.com', 1, 'Admin'),
(9, 'arifin', '1234', 'Arifin', 'arifin@example.com', 2, 'Admin'),
(10, 'fajar', '1234', 'fajar', 'fajar@example.com', 1, 'Admin');

-- --------------------------------------------------------

--
-- Struktur dari tabel `t_user_cabang`
--

DROP TABLE IF EXISTS `t_user_cabang`;
CREATE TABLE `t_user_cabang` (
  `cabang_id` int(20) UNSIGNED NOT NULL,
  `cabang` longtext DEFAULT NULL,
  `alias` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `t_user_cabang`
--

INSERT INTO `t_user_cabang` (`cabang_id`, `cabang`, `alias`) VALUES
(1, 'Jakarta', 'JKT'),
(2, 'Tangerang', 'TGR');

-- --------------------------------------------------------

--
-- Struktur dari tabel `t_user_group`
--

DROP TABLE IF EXISTS `t_user_group`;
CREATE TABLE `t_user_group` (
  `user_group_id` bigint(20) UNSIGNED NOT NULL,
  `user_group` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `t_user_group`
--

INSERT INTO `t_user_group` (`user_group_id`, `user_group`) VALUES
(1, 'Admin'),
(2, 'Super Admin'),
(4, 'Customer');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `t_user`
--
ALTER TABLE `t_user`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `t_user_cabang`
--
ALTER TABLE `t_user_cabang`
  ADD PRIMARY KEY (`cabang_id`) USING BTREE;

--
-- Indeks untuk tabel `t_user_group`
--
ALTER TABLE `t_user_group`
  ADD PRIMARY KEY (`user_group_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `t_user`
--
ALTER TABLE `t_user`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT untuk tabel `t_user_cabang`
--
ALTER TABLE `t_user_cabang`
  MODIFY `cabang_id` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `t_user_group`
--
ALTER TABLE `t_user_group`
  MODIFY `user_group_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
