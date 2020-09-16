package genesysapi_test

import (
	"encoding/json"
	"testing"

	"github.com/freman/genesysapi"
	"github.com/freman/genesysapi/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser(t *testing.T) {
	update := models.UpdateUser{}
	ob, err := json.Marshal(update)
	require.NoError(t, err)
	require.JSONEq(t, `{"acdAutoAnswer":false,"addresses":null,"certifications":null,"groups":null,"images":null,"locations":null,"primaryContactInfo":null,"profileSkills":null,"version":null}`, string(ob))

	update.Title = genesysapi.String("Ruler of the known universe")
	ob, err = json.Marshal(update)
	require.NoError(t, err)
	require.JSONEq(t, `{"acdAutoAnswer":false,"addresses":null,"certifications":null,"groups":null,"images":null,"locations":null,"primaryContactInfo":null,"profileSkills":null,"title":"Ruler of the known universe","version":null}`, string(ob))

	update.Title = genesysapi.String("")
	ob, err = json.Marshal(update)
	require.NoError(t, err)
	require.JSONEq(t, `{"acdAutoAnswer":false,"addresses":null,"certifications":null,"groups":null,"images":null,"locations":null,"primaryContactInfo":null,"profileSkills":null,"title":"","version":null}`, string(ob))

	update.Title = nil
	ob, err = json.Marshal(update)
	require.NoError(t, err)
	require.JSONEq(t, `{"acdAutoAnswer":false,"addresses":null,"certifications":null,"groups":null,"images":null,"locations":null,"primaryContactInfo":null,"profileSkills":null,"version":null}`, string(ob))
}
