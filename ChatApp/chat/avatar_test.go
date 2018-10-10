package main

import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
)

func TestAuthAvatar(t *testing.T){
	var authAvatar  AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL{
		t.Error("AuthAvatar.GetAvatarURL shour return ErrNoAvatarURL when no  value present")

	}
	//set value
	testUrl := "http://url-to-gavatar/"
	client.userData = map[string] interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarUrl shoudl return no error when value present")
	}

	if(url != testUrl){
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}


func testGravatarAvatar(t *testing.T){
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil{
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}

	if url != "https://www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346"{
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}


func testFileSystemAvatar(t *testing.T){
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename,[]byte{}, 0777)
	defer os.Remove(filename)
	var fileSystemAvatr FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"usedid":"abc"}
	url ,err := fileSystemAvatr.GetAvatarUrl(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg"{
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}
}