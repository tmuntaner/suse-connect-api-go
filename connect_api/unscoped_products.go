package connect_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UnscopedProduct struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Identifier       string      `json:"identifier"`
	FormerIdentifier string      `json:"former_identifier"`
	Version          string      `json:"version"`
	ReleaseType      interface{} `json:"release_type"`
	Arch             string      `json:"arch"`
	FriendlyName     string      `json:"friendly_name"`
	ProductClass     string      `json:"product_class"`
	Cpe              interface{} `json:"cpe"`
	Free             bool        `json:"free"`
	Description      interface{} `json:"description"`
	ReleaseStage     string      `json:"release_stage"`
	EulaURL          string      `json:"eula_url"`
	Repositories     []struct {
		ID               int    `json:"id"`
		URL              string `json:"url"`
		Name             string `json:"name"`
		DistroTarget     string `json:"distro_target"`
		Description      string `json:"description"`
		Enabled          bool   `json:"enabled"`
		Autorefresh      bool   `json:"autorefresh"`
		InstallerUpdates bool   `json:"installer_updates"`
	} `json:"repositories"`
	ProductType           string        `json:"product_type"`
	PredecessorIds        []interface{} `json:"predecessor_ids"`
	OnlinePredecessorIds  []interface{} `json:"online_predecessor_ids"`
	OfflinePredecessorIds []interface{} `json:"offline_predecessor_ids"`
	Shortname             interface{}   `json:"shortname"`
	Recommended           bool          `json:"recommended"`
	Extensions            []struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		Identifier       string      `json:"identifier"`
		FormerIdentifier string      `json:"former_identifier"`
		Version          string      `json:"version"`
		ReleaseType      interface{} `json:"release_type"`
		Arch             string      `json:"arch"`
		FriendlyName     string      `json:"friendly_name"`
		ProductClass     string      `json:"product_class"`
		Cpe              interface{} `json:"cpe"`
		Free             bool        `json:"free"`
		Description      interface{} `json:"description"`
		ReleaseStage     string      `json:"release_stage"`
		EulaURL          string      `json:"eula_url"`
		Repositories     []struct {
			ID               int    `json:"id"`
			URL              string `json:"url"`
			Name             string `json:"name"`
			DistroTarget     string `json:"distro_target"`
			Description      string `json:"description"`
			Enabled          bool   `json:"enabled"`
			Autorefresh      bool   `json:"autorefresh"`
			InstallerUpdates bool   `json:"installer_updates"`
		} `json:"repositories"`
		ProductType           string        `json:"product_type"`
		PredecessorIds        []interface{} `json:"predecessor_ids"`
		OnlinePredecessorIds  []interface{} `json:"online_predecessor_ids"`
		OfflinePredecessorIds []interface{} `json:"offline_predecessor_ids"`
		Shortname             interface{}   `json:"shortname"`
		Recommended           bool          `json:"recommended"`
		Extensions            []interface{} `json:"extensions"`
	} `json:"extensions"`
}

func (s *Service) UnscopedProducts() ([]UnscopedProduct, error) {

	url := "https://scc.suse.com/connect/organizations/products/unscoped?page=1"
	var products []UnscopedProduct
	var err error

	for {
		next, unscoped, err := apiUnscopedProducts(url, s.Username, s.Password)
		products = append(products, unscoped...)

		if err != nil {
			return nil, err
		}

		if next == "" {
			break
		}

		url = next
	}

	return products, err
}

func apiUnscopedProducts(url string, username string, password string) (string, []UnscopedProduct, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Accept", "application/vnd.scc.suse.com.v4+json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", nil, err
	}

	var allProducts []UnscopedProduct
	body := readBody(resp.Body)
	err = json.Unmarshal(body, &allProducts)
	if err != nil {
		fmt.Println(string(body))
		return "", nil, err
	}

	linkHeader := resp.Header.Get("link")
	next := getNext(linkHeader)

	return next, allProducts, nil
}
