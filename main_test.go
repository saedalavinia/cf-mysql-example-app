package main_test

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/EngineerBetter/cf-mysql-example-app"
)

var _ = Describe("The App", func() {
	It("Does a thing", func() {
		repo := NewInMemoryRepository()
		handler := NewPutGetHandler(repo)
		server := httptest.NewServer(handler)
		defer server.Close()
		client := http.DefaultClient
		request, err := http.NewRequest(http.MethodPut, server.URL+"/somevalue", strings.NewReader("testvalue"))
		Ω(err).ShouldNot(HaveOccurred())
		response, err := client.Do(request)
		Ω(err).ShouldNot(HaveOccurred())
		buff := new(bytes.Buffer)
		buff.ReadFrom(response.Body)
		responseBody := buff.String()
		Ω(responseBody).Should(Equal("created"))

		response, err = http.Get(server.URL + "/somevalue")
		Ω(err).ShouldNot(HaveOccurred())
		buff = new(bytes.Buffer)
		buff.ReadFrom(response.Body)
		responseBody = buff.String()
		Ω(responseBody).Should(Equal("testvalue"))
	})
})

type InMemoryRepository struct {
	KeyValues map[string]string
}

func NewInMemoryRepository() InMemoryRepository {
	var r InMemoryRepository
	r.KeyValues = make(map[string]string)
	return r
}

func (r InMemoryRepository) Write(key, value string) error {
	r.KeyValues[key] = value
	return nil
}

func (r InMemoryRepository) Read(key string) (string, error) {
	return r.KeyValues[key], nil
}
