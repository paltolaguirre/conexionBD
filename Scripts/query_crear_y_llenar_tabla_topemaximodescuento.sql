DO $$
 DECLARE
    _kind "char";
BEGIN

 SELECT relkind
		   FROM   pg_class
		   WHERE  relname = 'topmaximodescuento_id_seq'  -- sequence name
		   INTO  _kind;
		   IF NOT FOUND THEN       -- name is free
		      CREATE SEQUENCE topmaximodescuento_id_seq
				  INCREMENT 1
				  MINVALUE 1
				  MAXVALUE 9223372036854775807
				  START 1
				  CACHE 1;
				ALTER TABLE topmaximodescuento_id_seq
				  OWNER TO postgres;
				GRANT ALL ON SEQUENCE topmaximodescuento_id_seq TO postgres;
				GRANT SELECT, USAGE ON SEQUENCE topmaximodescuento_id_seq TO rol_super;
				GRANT SELECT, USAGE ON SEQUENCE topmaximodescuento_id_seq TO rol_common;
		   ELSIF _kind = 'S' THEN  -- sequence exists
		      -- nada
		   ELSE                    -- object name exists for different kind
		      -- nada
		   END IF;

CREATE TABLE topemaximodescuento(
   id INTEGER NOT NULL DEFAULT nextval('topmaximodescuento_id_seq'::regclass),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  nombre VARCHAR,
  codigo VARCHAR,
  descripcion VARCHAR,
  activo int,
  topecasomuerte NUMERIC(19,4) NOT NULL,
  topeseguroahorro NUMERIC(19,4) NOT NULL,
  toperetiroprivado NUMERIC(19,4) NOT NULL,
  topesepelio NUMERIC(19,4) NOT NULL,
  topehipotecarios NUMERIC(19,4) NOT NULL,
  anio INT UNIQUE,
  CONSTRAINT topmaximodescuento_seq PRIMARY KEY (id)
)

WITH (
	OIDS=FALSE
);
	
END;
$$;



DO $$
BEGIN

INSERT INTO topemaximodescuento(created_at,activo,topecasomuerte,topeseguroahorro,toperetiroprivado,topesepelio,topehipotecarios,anio)
VALUES
(current_timestamp, 1, 12000.00, 12000.00, 12000.00, 993.26, 20000.00, 2019),
(current_timestamp, 1, 18000.00, 18000.00, 18000.00, 993.26, 20000.00, 2020);

UPDATE public.topemaximodescuento SET id = -id;

END;
$$;