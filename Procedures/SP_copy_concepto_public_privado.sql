CREATE OR REPLACE FUNCTION ST_copy_concepto_public_privado() RETURNS void AS $$
BEGIN
        insert into concepto select * from public.concepto where public.concepto.id not in (select id from concepto);

        EXECUTE (
            SELECT
                                'UPDATE concepto as v SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.concepto.' || quote_ident(column_name), ',') || ') FROM public.concepto WHERE v.id = public.concepto.id'
            FROM   information_schema.columns
            WHERE  table_name   = 'concepto'       -- table name, case sensitive
              AND    table_schema = 'public'  -- schema name, case sensitive
              AND    column_name <> 'id'      -- all columns except id
        );

END
$$ LANGUAGE plpgsql;