-- Migration script for Evermos database
-- This script creates the database schema by zaki

-- Create database
CREATE DATABASE IF NOT EXISTS evermos CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE evermos;

-- Table structure for table `users`
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kata_sandi` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `notelp` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `jenis_kelamin` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tentang` text COLLATE utf8mb4_general_ci,
  `pekerjaan` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `id_provinsi` int DEFAULT NULL,
  `id_kota` int DEFAULT NULL,
  `isAdmin` tinyint(1) DEFAULT '0',
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `notelp` (`notelp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `toko`
CREATE TABLE IF NOT EXISTS `toko` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `nama_toko` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `url_toko` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `toko_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `alamat`
CREATE TABLE IF NOT EXISTS `alamat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `judul_alamat` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `nama_penerima` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `no_telp` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `detail_alamat` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `alamat_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `category`
CREATE TABLE IF NOT EXISTS `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_category` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `produk`
CREATE TABLE IF NOT EXISTS `produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_produk` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `slug` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `harga_reseller` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `harga_konsumen` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `stok` int DEFAULT NULL,
  `deskripsi` text COLLATE utf8mb4_general_ci,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `id_toko` int DEFAULT NULL,
  `id_category` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_toko` (`id_toko`),
  KEY `id_category` (`id_category`),
  CONSTRAINT `produk_ibfk_1` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`),
  CONSTRAINT `produk_ibfk_2` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `foto_produk`
CREATE TABLE IF NOT EXISTS `foto_produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_produk` int DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_produk` (`id_produk`),
  CONSTRAINT `foto_produk_ibfk_1` FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `log_produk`
CREATE TABLE IF NOT EXISTS `log_produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_produk` int DEFAULT NULL,
  `nama_produk` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `slug` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `harga_reseller` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `harga_konsumen` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `deskripsi` text COLLATE utf8mb4_general_ci,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `id_toko` int DEFAULT NULL,
  `id_category` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_produk` (`id_produk`),
  KEY `id_toko` (`id_toko`),
  KEY `id_category` (`id_category`),
  CONSTRAINT `log_produk_ibfk_1` FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id`),
  CONSTRAINT `log_produk_ibfk_2` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`),
  CONSTRAINT `log_produk_ibfk_3` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `trx`
CREATE TABLE IF NOT EXISTS `trx` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `alamat_pengiriman` int DEFAULT NULL,
  `harga_total` int DEFAULT NULL,
  `kode_invoice` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `method_bayar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `alamat_pengiriman` (`alamat_pengiriman`),
  CONSTRAINT `trx_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`),
  CONSTRAINT `trx_ibfk_2` FOREIGN KEY (`alamat_pengiriman`) REFERENCES `alamat` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Table structure for table `detail_trx`
CREATE TABLE IF NOT EXISTS `detail_trx` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_trx` int DEFAULT NULL,
  `id_log_produk` int DEFAULT NULL,
  `id_toko` int NOT NULL,
  `kuantitas` int DEFAULT NULL,
  `harga_total` int DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_tx` (`id_trx`),
  KEY `id_log_produk` (`id_log_produk`),
  KEY `detail_trx_ibfk_3` (`id_toko`),
  CONSTRAINT `detail_trx_ibfk_1` FOREIGN KEY (`id_trx`) REFERENCES `trx` (`id`),
  CONSTRAINT `detail_trx_ibfk_2` FOREIGN KEY (`id_log_produk`) REFERENCES `log_produk` (`id`),
  CONSTRAINT `detail_trx_ibfk_3` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
