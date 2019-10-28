INSERT INTO tipoliquidacion(created_at, nombre, codigo, descripcion, activo)
            values
(current_timestamp,'Importe Remunerativo','IMPORTE_REMUNERATIVO','',1),
(current_timestamp,'Importe No Remunerativo','IMPORTE_NO_REMUNERATIVO','',1),
(current_timestamp,'Descuento','DESCUENTO','',1),
(current_timestamp,'Retención','RETENCION','',1),
(current_timestamp,'Aporte Patronal','APORTE_PATRONAL','',1);

UPDATE public.estadocivil SET id = -id;