### Go test

```
go test -v
```

### Gen mock from interface

```
mockery -name=UserRepo
mockery -all
```

### Coverage

```
Given a coverage profile produced by 'go test':
        go test -coverprofile=c.out

Open a web browser displaying annotated source code:
        go tool cover -html=c.out

Write out an HTML file instead of launching a web browser:
        go tool cover -html=c.out -o coverage.html

Display coverage percentages to stdout for each function:
        go tool cover -func=c.out

Finally, to generate modified source code with coverage annotations
(what go test -cover does):
        go tool cover -mode=set -var=CoverageVariableName program.go

Flags:
  -V    print version and exit
  -func string
        output coverage profile information for each function
  -html string
        generate HTML representation of coverage profile
  -mode string
        coverage mode: set, count, atomic
  -o string
        file for output; default: stdout
  -var string
        name of coverage variable to generate (default "GoCover")

  Only one of -html, -func, or -mode may be set.
```

### Example with test mongo

```
func TestUserRepoImpl_Insert(t *testing.T) {
	cs := testUtils.ConnString(t)
	connector := driverMongo.ConnectMongoDB(cs.Original)
	db := connector.Client.Database(testUtils.GetDBName(cs))

	assert.Equal(t, "mongodb://localhost:27018/go_learning_test", cs.Original)

	defer func() {
		// _ = db.Drop(context.Background())
		_ = db.Collection(colName).Drop(context.Background())
		_ = connector.Client.Disconnect(context.Background())
	}()

	userRepo := NewUserRepo(db)
	...
}
```