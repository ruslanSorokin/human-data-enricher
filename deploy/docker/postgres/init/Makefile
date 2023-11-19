POSTGRES_HOST=localhost
POSTGRES_PORT=5432

init:
	@psql -U postgres -d postgres -a -f              000001_create-database.up.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000002_create-schema.up.sql                  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d postgres -a -f              000003_create-role-ddl.up.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000004_assign-privileges-to-role-ddl.up.sql  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d postgres -a -f              000005_create-role-dml.up.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000006_assign-privileges-to-role-dml.up.sql  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000007_create-user-migrator.up.sql           --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000008_create-user-application.up.sql        --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U migrator -d 'human-data-enricher' -a -f 000009_assign-default-privileges.up.sql      --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}



deinit:
	@psql -U migrator -d 'human-data-enricher' -a -f 000009_assign-default-privileges.down.sql      --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000008_create-user-application.down.sql        --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000007_create-user-migrator.down.sql           --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000006_assign-privileges-to-role-dml.down.sql  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d postgres -a -f              000005_create-role-dml.down.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000004_assign-privileges-to-role-ddl.down.sql  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d postgres -a -f              000003_create-role-ddl.down.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d 'human-data-enricher' -a -f 000002_create-schema.down.sql                  --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
	@psql -U postgres -d postgres -a -f              000001_create-database.down.sql                --host ${POSTGRES_HOST} --port ${POSTGRES_PORT}
