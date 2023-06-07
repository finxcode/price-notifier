DROP TABLE IF EXISTS basic_divergences;
CREATE TABLE IF NOT EXISTS basic_divergences (
    id INT NOT NULL AUTO_INCREMENT,
    coin_id INT NOT NULL,
    baseline_id INT NOT NULL,
    symbol VARCHAR(300) NOT NULL,
    divergence24H FLOAT NOT NULL,
    divergence7D FLOAT NOT NULL,
    divergence_total FLOAT NOT NULL,
    trading_day VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;