package nutil

import "testing"

func TestConfig(t *testing.T) {

	t.Run("Default Config Test", func(t *testing.T) {

		cfg, err := CreateNewConfigFromPath("test.conf")
		if err != nil {
			t.Fatal(err)
		}

		cls := cfg.GetArrayConfigObject("cluster")
		if cls == nil {
			t.Fatal("Arr Convert Fail")
		}

	})

}
