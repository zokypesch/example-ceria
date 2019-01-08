package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllConnection(t *testing.T) {

	t.Run("Test Get DB", func(t *testing.T) {
		conn := GetDB()
		db, errDB := conn.GetConn()
		defer db.Close()
		assert.NoError(t, errDB)
	})

	t.Run("Test Get Router", func(t *testing.T) {
		ginCfg, els := GetRouter()

		_, errGin := ginCfg.Register(false)
		assert.NoError(t, errGin)
		assert.NotNil(t, els.Status)
	})

	t.Run("Tes get group nil", func(t *testing.T) {
		grp := GetGroup("")

		assert.Empty(t, grp.Name)
	})

	t.Run("Tes get group", func(t *testing.T) {
		grp := GetGroup("tesst")

		assert.NotEmpty(t, grp.Name)
	})

}
