package calculations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetAmount(t *testing.T) {
	tests := []struct {
		name      string
		parameter string
		want      float64
		err       string
	}{
		{
			name:      "should convert a valid string to float ignoring positive symbol",
			parameter: "+10",
			want:      10,
		},
		{
			name:      "should convert a valid string to float ignoring negative symbol",
			parameter: "-20",
			want:      20,
		},
		{
			name:      "should show an error trying to convert asd to float",
			parameter: "asd",
			want:      10,
			err:       "cannot convert amount asd: strconv.ParseFloat: parsing \"sd\": invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := GetAmount(tt.parameter)
			if tt.err != "" {
				assert.Equal(t, tt.err, err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, amount)
			}
		})
	}
}
