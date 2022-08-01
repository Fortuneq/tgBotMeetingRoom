POSTGRESQL_URL='postgres://postgres:537j04222@localhost:5432/postgres?sslmode=disable'

upmigrate:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up
downmigrate:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down