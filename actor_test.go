package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/BlackAWhite8/vkTest/actors"
)

func TestGetActors(t *testing.T) {
	t.Run("get actors test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/actors", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		actors.GetActors(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}

func TestDeleteActor(t *testing.T) {
	t.Run("delete actor test", func(t *testing.T) {
		b := []byte(
			`{
				"name":"Johny Depp",
				"sex":"male",
				"birthday":"1977-04-24"
			}`)
		r := bytes.NewReader(b)
		req, err := http.NewRequest("DELETE", "/actors/delete", r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		actors.DeleteActor(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}

func TestCreateActor(t *testing.T) {
	t.Run("create actor test", func(t *testing.T) {
		b := []byte(
			`{
				"name":"Lucy Lee",
				"sex" : "female",
				"birthday": "1978-01-01"
			}`)
		r := bytes.NewReader(b)
		req, err := http.NewRequest("POST", "/actors/create", r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		actors.CreateActor(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}

func TestUpdateActor(t *testing.T) {
	t.Run("update movies test", func(t *testing.T) {
		/* изменений не должно быть если такого фильма нет*/
		b := []byte(
			`{
				"name":"Johny_Depp",
				"sex" : "m",
				"birthday":"1977-04-24"
			}`)
		r := bytes.NewReader(b)
		v := url.Values{}
		v.Add("name", "Johny Depp")
		v.Add("sex", "male")
		req, err := http.NewRequest("PUT", "/actors/update?"+v.Encode(), r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		actors.UpdateActor(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}
