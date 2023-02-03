ALTER TABLE novedades.novedades_poscontractuales ADD COLUMN estado VARCHAR(15);

UPDATE TABLE novedades.novedades_poscontractuales SET estado='TERMINADA';