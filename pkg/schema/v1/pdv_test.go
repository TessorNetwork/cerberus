package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/TessorNetwork/cerberus/pkg/schema/types"
)

func TestPDV_Validate(t *testing.T) {
	require.True(t, PDV{
		Cookie{
			Timestamp: types.Timestamp{Time: time.Now()},
			Source:    types.Source{Host: "https://furya.xyz"},
			Name:      "cookie",
			Value:     "value",
		},
		Profile{
			FirstName: "First",
			LastName:  "Last",
			Emails:    []string{"test@furya.xyz"},
			Gender:    types.GenderMale,
			Avatar:    "https://furya.xyz/avatar.jpeg",
			Birthday:  mustDate("1990-01-01"),
		},
	}.Validate())
}

func TestPDV_Validate_invalid(t *testing.T) {
	require.False(t, PDV{}.Validate())

	require.False(t, PDV{
		Cookie{
			Source: types.Source{Host: "https://furya.xyz"},
			Name:   "cookie",
		},
	}.Validate())
}

func mustDate(s string) *types.Date {
	var d types.Date

	if err := d.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}

	return &d
}
