ALTER TABLE novedades.novedades_poscontractuales
DROP COLUMN IF EXISTS estado,
DROP COLUMN IF EXISTS enlace_documento,
DROP COLUMN IF EXISTS oficio_supervisor,
DROP COLUMN IF EXISTS oficio_ordenador;
