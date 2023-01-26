package main

import (
	"context"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	dbPath := "projects/my-project/instance/my-instance/databases/my-database"
	client, err := spanner.NewClient(ctx, dbPath)
	if err != nil {
		log.Fatal(err)
	}
	iter := client.Single().Query(
		ctx,
		spanner.Statement{
			`SELECT tableA.config_name, tableA.policy,
					FROM tableA, tableB
					WHERE tableA.config_name = tableB.config_name
					AND tableB.owner = "name"
					ORDER BY tableA.config_name ASC`,
			map[string]interface{}{},
		})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			log.Println("Done reading rows")
			return
		}
		if err != nil {
			panic(err)
		}
		var configName string
		var policy []byte
		if err := row.Columns(&configName, &policy); err != nil {
			panic(err)
		}
	}
}
