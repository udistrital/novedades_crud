ALTER TABLE novedades.novedades_poscontractuales ADD COLUMN enlace_documento VARCHAR(100);

UPDATE novedades.novedades_poscontractuales SET enlace_documento='';