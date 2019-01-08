package migration

import (
	"testing"

	"github.com/zokypesch/ceria/core"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	migrate := NewListMigration()

	assert.NoError(migrate.Migrate(core.GetTestConnection()))
	// error because fake connection gorm doesent have a CReate Function
}
