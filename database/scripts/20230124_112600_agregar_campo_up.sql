ALTER TABLE novedades.novedades_poscontractuales
ADD COLUMN IF NOT EXISTS enlace_documento VARCHAR(30)
ADD COLUMN IF NOT EXISTS oficio_supervisor VARCHAR(15),
ADD COLUMN IF NOT EXISTS oficio_ordenador VARCHAR(15),
ADD COLUMN IF NOT EXISTS estado VARCHAR(15);

UPDATE novedades.novedades_poscontractuales
SET estado='4519', oficio_supervisor='', oficio_ordenador='';