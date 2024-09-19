createdb:
	docker exec -it postgresql_plat createdb --username=root --owner=root platform_transactions	 
dropdb:
	docker exec -it postgresql_plat dropdb  platform_transactions
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/platform_transactions?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/platform_transactions?sslmode=disable" -verbose down
sqlc: 
	sqlc generate		 	
.PHONY: createdb dropdb migrateup migratedown sqlc