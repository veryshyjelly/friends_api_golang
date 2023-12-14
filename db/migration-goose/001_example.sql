-- +goose Up

CREATE TABLE `accounts` (
  `id` bigint(10) PRIMARY KEY,
  `owner` varchar(10) NOT NULL,
  `balance` bigint(10) NOT NULL,
  `currency` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT now());
