SELECT declaraciones_juradas.id,
tipos_caracter_declaracion.descripcion as 'caracter_declaracion',
declaraciones_juradas.version,
declaraciones_juradas.tipo_estado,
declaraciones_juradas.created_at,
declaraciones_juradas.accepted_at,
declaraciones_juradas.received_at,
declaraciones_juradas.deleted_at,
personas.nombres,
personas.apellido,
personas.documento,
fuerzas.descripcion as 'fuerza'
FROM declaraciones_juradas
INNER JOIN personas on personas.id = declaraciones_juradas.id_persona
INNER JOIN fuerzas on fuerzas.id = personas.id_fuerza
INNER JOIN tipos_caracter_declaracion on declaraciones_juradas.id_tipo_caracter = tipos_caracter_declaracion.id
WHERE 1=1
AND personas.id_fuerza = 2
AND declaraciones_juradas.created_at <= 'NOW()'
AND declaraciones_juradas.tipo_estado IN ('Entregado en Dependencia','Recibido','Entregado a RRHH')