DROP FUNCTION IF EXISTS ST_eliminar_liquidaciones_sin_legajo();

CREATE OR REPLACE FUNCTION ST_eliminar_liquidaciones_sin_legajo() RETURNS void AS $$
DECLARE
    schema RECORD;
BEGIN
    FOR schema IN SELECT schema_name FROM information_schema.schemata INNER JOIN information_schema.tables ON table_schema = schema_name WHERE left(schema_name, 3) = 'tnt' AND table_name = 'liquidacion'
        LOOP
            EXECUTE format('delete from %I.liquidacion where legajoid is null', schema.schema_name);
        END LOOP;
END;
$$ LANGUAGE plpgsql;


select * from ST_eliminar_liquidaciones_sin_legajo();

DROP FUNCTION IF EXISTS ST_eliminar_liquidaciones_sin_legajo();