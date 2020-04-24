CREATE OR REPLACE FUNCTION ST_copy_formulas_public_privado() RETURNS void AS $$
BEGIN
        insert into value select * from public.value where public.value.id not in (select id from value);
        insert into function select * from public.function where public.function.name not in (select name from function);
        insert into param select * from public.param where public.param.id not in (select id from param);

        EXECUTE (
            SELECT
                                'UPDATE value as v SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.value.' || quote_ident(column_name), ',') || ') FROM public.value WHERE v.id = public.value.id'
            FROM   information_schema.columns
            WHERE  table_name   = 'value'       -- table name, case sensitive
              AND    table_schema = 'public'  -- schema name, case sensitive
              AND    column_name <> 'id'      -- all columns except id
        );

        EXECUTE (
            SELECT
                                'UPDATE function as fun SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.function.' || quote_ident(column_name), ',') || ') FROM public.function WHERE fun.name = public.function.name'
            FROM   information_schema.columns
            WHERE  table_name   = 'function'       -- table name, case sensitive
              AND    table_schema = 'public'  -- schema name, case sensitive
              AND    column_name <> 'name'      -- all columns except name
        );

        EXECUTE (
            SELECT
                                'UPDATE param as par SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.param.' || quote_ident(column_name), ',') || ') FROM public.param WHERE par.id = public.param.id'
            FROM   information_schema.columns
            WHERE  table_name   = 'param'       -- table name, case sensitive
              AND    table_schema = 'public'  -- schema name, case sensitive
              AND    column_name <> 'id'      -- all columns except id
        );
END
$$ LANGUAGE plpgsql;