DO $$
 DECLARE
    _kind "char";
BEGIN

    SELECT relkind
    FROM   pg_class
    WHERE  relname = 'deduccionespersonales_id_seq'  -- sequence name
    INTO  _kind;
    IF NOT FOUND THEN       -- name is free
        CREATE SEQUENCE deduccionespersonales_id_seq
            INCREMENT 1
            MINVALUE 1
            MAXVALUE 9223372036854775807
            START 1
            CACHE 1;
        ALTER TABLE deduccionespersonales_id_seq
            OWNER TO postgres;
        GRANT ALL ON SEQUENCE deduccionespersonales_id_seq TO postgres;
        GRANT SELECT, USAGE ON SEQUENCE deduccionespersonales_id_seq TO rol_super;
        GRANT SELECT, USAGE ON SEQUENCE deduccionespersonales_id_seq TO rol_common;
    ELSIF _kind = 'S' THEN  -- sequence exists
        -- nada
    ELSE                    -- object name exists for different kind
        -- nada
    END IF;


CREATE TABLE deduccionespersonales(
  id INTEGER NOT NULL DEFAULT nextval('deduccionespersonales_id_seq'::regclass),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  nombre VARCHAR,
  codigo VARCHAR,
  descripcion VARCHAR,
  activo int,
  valorfijoconyuge NUMERIC(19,4) NOT NULL,
  valorfijohijo NUMERIC(19,4) NOT NULL,
  valorfijomni NUMERIC(19,4) NOT NULL,
  valorfijoddei NUMERIC(19,4) NOT NULL,
  anio INT UNIQUE,
   CONSTRAINT deduccionespersonales_seq PRIMARY KEY (id)
)

WITH (
	OIDS=FALSE
);
	
END;
$$;



$$
BEGIN
    INSERT INTO deduccionespersonales(created_at,activo,valorfijoconyuge,valorfijohijo,valorfijomni,valorfijoddei,anio)
    VALUES
    (current_timestamp, 1, 80033.90, 40361.43, 103018.79, 494490.18, 2019),
    (current_timestamp, 1, 115471.38, 58232.65, 123861.17, 594533.62, 2020);

    UPDATE public.deduccionespersonales SET id = -id;
END;
$$;