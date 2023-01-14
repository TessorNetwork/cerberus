package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/TessorNetwork/cerberus/pkg/schema/types"
)

func TestSearchHistory_Validate(t *testing.T) {
	tt := []struct {
		name  string
		d     SearchHistory
		valid bool
	}{
		{
			name: "valid",
			d: SearchHistory{
				Timestamp: types.Timestamp{Time: time.Now()},
				Engine:    "furya",
				Domain:    "furya.xyz",
				Query:     "the best crypto",
			},
			valid: true,
		},
		{
			name: "empty engine",
			d: SearchHistory{
				Timestamp: types.Timestamp{Time: time.Now()},
				Engine:    "",
				Domain:    "furya.xyz",
				Query:     "the best crypto",
			},
			valid: false,
		},
		{
			name: "empty searchLine",
			d: SearchHistory{
				Timestamp: types.Timestamp{Time: time.Now()},
				Engine:    "furya",
				Domain:    "furya.xyz",
				Query:     "",
			},
			valid: false,
		},
		{
			name: "empty domain",
			d: SearchHistory{
				Timestamp: types.Timestamp{Time: time.Now()},
				Engine:    "furya",
				Domain:    "",
				Query:     "",
			},
			valid: false,
		},
		{
			name: "invalid timestamp",
			d: SearchHistory{
				Engine: "furya",
				Query:  "something",
				Domain: "furya.xyz",
			},
			valid: false,
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.valid, tc.d.Validate())
		})
	}
}
