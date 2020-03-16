drop view if exists liquidaciones_all;

DROP FUNCTION IF EXISTS ST_liquidaciones_tenant();

CREATE OR REPLACE FUNCTION ST_liquidaciones_tenant() RETURNS table (liquidacionid integer, created_at timestamp, tenant varchar) AS $$
DECLARE
    schema RECORD;
BEGIN
    FOR schema IN SELECT schema_name FROM information_schema.schemata INNER JOIN information_schema.tables ON table_schema = schema_name WHERE left(schema_name, 4) = 'tnt_' AND table_name = 'liquidacion'
        LOOP
            RETURN QUERY EXECUTE
                format('SELECT id as liquidacionid, created_at::timestamp as created_at, ''%s''::VARCHAR as tenant FROM %I.liquidacion', schema.schema_name, schema.schema_name);
        END LOOP;
END;
$$ LANGUAGE plpgsql;

create view liquidaciones_all as select * from ST_liquidaciones_tenant() order by created_at DESC;

GRANT ALL ON FUNCTION ST_liquidaciones_tenant to read_only;
GRANT ALL ON liquidaciones_all to read_only;
