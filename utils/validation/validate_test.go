package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUrlPicture(t *testing.T) {
	t.Run("valid-param-url", func(t *testing.T) {
		var url_valid []string = []string{"http://image.png", "http://image.jpeg", "https://image.jpg"}

		for _, val := range url_valid {
			assert.Nil(t, ValideteUrlPicture(val))
		}
	})

	t.Run("invalid-param-url", func(t *testing.T) {
		var url_invalid []string = []string{"http://image.pnge", "https://image.jp1eg", "htt1ps://image.jpg", "gha_://image.png"}

		for _, val := range url_invalid {
			assert.Error(t, ValideteUrlPicture(val))
		}
	})

}

func TestValideteText(t *testing.T) {
	t.Run("valid-param-text", func(t *testing.T) {
		text := "erhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdv"
		assert.Nil(t, ValideteText(text))
	})

	t.Run("invalid-param-text", func(t *testing.T) {
		text := "kvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdverhjgbdbaskasdbsbdjvbasdkvbksadvbkasbdlvkbasdvblasdkvblksbdvkbasdvbasdvbkasdlvkasdlvasldkvasdv"
		assert.Error(t, ValideteText(text))
	})
}

func TestValidateUrlQuery(t *testing.T) {
	t.Run("valid-param-query", func(t *testing.T) {
		var limit_valid []string = []string{"0", "1", "2", "3", "4", "5", "6"}
		var offset_valid []string = []string{"0", "1", "2", "3", "4", "5", "6"}
		for i := 0; i < len(limit_valid); i++ {
			assert.Nil(t, ValidateUrlQuery(limit_valid[i], offset_valid[i], date, asc))
		}
	})

	t.Run("invalid-param-query", func(t *testing.T) {
		var limit_valid []string = []string{"", "g", "2aw", "3jn", "4xx", "_", ""}
		var offset_valid []string = []string{"0d", "1ad", "2__1", "3hh", "411t", "522222av", "6xcb"}
		for i := 0; i < len(limit_valid); i++ {
			assert.Error(t, ValidateUrlQuery(limit_valid[i], offset_valid[i], date, asc))
		}
	})

}
