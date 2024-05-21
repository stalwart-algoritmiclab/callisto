package vault

import (
	"os"
	"testing"
)

func TestPush(t *testing.T) {
	client, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	if err = client.Push("rate-parser", "api:\n\thost: Nats"); err != nil {
		t.Fatal(err)
	}

	data, _ := client.Pull("rate-parser")
	t.Log(string(data))
}

func TestPushRaw(t *testing.T) {
	client, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	if err = client.PushRaw("test", map[string]any{
		"first secret":  "some description",
		"second secret": "some description",
	}); err != nil {
		t.Fatal(err)
	}

	data, _ := client.PullRaw("test")
	t.Log(data)
}

func TestPull(t *testing.T) {
	client, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	data, err := client.Pull("rate-parser")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestPullRaw(t *testing.T) {
	client, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	data, err := client.PullRaw("system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}

func TestDownload(t *testing.T) {
	client, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	if err = client.Download("rate-parser", "config.yaml"); err != nil {
		t.Error(err)
	}

	_, err = os.Open("config.yaml")
	if !os.IsNotExist(err) {
		os.Remove("config.yaml")
		return
	}

	t.Error("file don`t download")
}

func TestUpload(t *testing.T) {
	cl, err := NewClient(NamespaceCubbyhole)
	if err != nil {
		t.Fatal(err)
	}

	if err = cl.Upload("rate-parser", "config_test.yaml"); err != nil {
		t.Fatal(err)
	}

	data, err := cl.Pull("rate-parser")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
