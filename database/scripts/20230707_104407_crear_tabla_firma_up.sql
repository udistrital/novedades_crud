
CREATE TABLE IF NOT EXISTS novedades.aprobacionFirma (
    id serial not null primary key,
    proceso varchar(20) not null,
    fecha_proceso timestamp not null,
    documento_persona int not null,
    nombre_persona varchar(50),
    documento_acta text not null,
    fecha_creacion timestamp,
    fecha_modificacion timestamp,
    activo boolean,
    id_novedades_poscontractuales Integer not null
);

ALTER TABLE novedades.aprobacionFirma
ADD CONSTRAINT novedades_poscontractuales_fk
FOREIGN KEY (id_novedades_poscontractuales)
REFERENCES novedades.novedades_poscontractuales (id);