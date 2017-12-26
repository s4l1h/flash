package flash

import "testing"

func Test(t *testing.T) {

	message := "işlem başarılı"

	msg := New()
	msg.AddSuccess(message)
	res := msg.GetSuccess()
	if res[0].Data != message {
		t.Error(res)
	}
	if res[0].Type != "success" {
		t.Error(res)
	}
}
