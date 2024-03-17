package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/BlackAWhite8/vkTest/movies"
)

func TestGetMovies(t *testing.T) {
	t.Run("get movies test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/movies", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		movies.GetMovies(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}

func TestSearchMovies(t *testing.T) {
	t.Run("search movies test", func(t *testing.T) {
		/*search with param title*/
		v := url.Values{}
		v.Add("search", "title")
		req, err := http.NewRequest("GET", "/movies/search?"+v.Encode(), nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		movies.SearchMovie(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())

		/* test with default param (score)*/
		req2, err := http.NewRequest("GET", "/movies/search", nil)
		if err != nil {
			t.Fatal(err)
		}

		res2 := httptest.NewRecorder()
		movies.SearchMovie(res2, req2)

		fmt.Println(res2.Result())
		fmt.Println(res2.Body.String())
	})
}

func TestCreateMovie(t *testing.T) {
	t.Run("create movies test", func(t *testing.T) {
		b := []byte(
			`{
				"title":"rancho",
				"description":"desc",
				"realise_date":"1899-01-01",
				"score": 0
			}`)
		r := bytes.NewReader(b)
		req, err := http.NewRequest("POST", "/movies/create", r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		movies.CreateMovie(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())

	})
}

func TestDeleteMovie(t *testing.T) {
	t.Run("delete movies test", func(t *testing.T) {
		b := []byte(
			`{
				"title":"a",
				"description":"test 2",
				"realise_date":"1977-08-07",
				"score": 10
			}`)
		r := bytes.NewReader(b)
		req, err := http.NewRequest("DELETE", "/movies/delete", r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		movies.DeleteMovie(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}

func TestUpdateMovie(t *testing.T) {
	t.Run("update movies test", func(t *testing.T) {
		/* изменений не должно быть если такого фильма нет*/
		b := []byte(
			`{
				"title":"a",
				"description":"test 2",
				"realise_date":"1977-08-07",
				"score": 10
			}`)
		r := bytes.NewReader(b)
		v := url.Values{}
		v.Add("title", "rancho")
		req, err := http.NewRequest("PUT", "/movies/update?"+v.Encode(), r)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		movies.UpdateMovie(res, req)

		fmt.Println(res.Result())
		fmt.Println(res.Body.String())
	})
}
