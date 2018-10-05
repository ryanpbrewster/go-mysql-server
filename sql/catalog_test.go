package sql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/src-d/go-mysql-server.v0/mem"
	"gopkg.in/src-d/go-mysql-server.v0/sql"
)

func TestCatalogDatabase(t *testing.T) {
	require := require.New(t)

	c := sql.NewCatalog()
	db, err := c.Database("foo")
	require.EqualError(err, "database not found: foo")
	require.Nil(db)

	mydb := mem.NewDatabase("foo")
	c.AddDatabase(mydb)

	db, err = c.Database("foo")
	require.NoError(err)
	require.Equal(mydb, db)
}

func TestCatalogTable(t *testing.T) {
	require := require.New(t)

	c := sql.NewCatalog()

	table, err := c.Table("foo", "bar")
	require.EqualError(err, "database not found: foo")
	require.Nil(table)

	db := mem.NewDatabase("foo")
	c.AddDatabase(db)

	table, err = c.Table("foo", "bar")
	require.EqualError(err, "table not found: bar")
	require.Nil(table)

	mytable := mem.NewTable("bar", nil)
	db.AddTable("bar", mytable)

	table, err = c.Table("foo", "bar")
	require.NoError(err)
	require.Equal(mytable, table)

	table, err = c.Table("foo", "BAR")
	require.NoError(err)
	require.Equal(mytable, table)
}
