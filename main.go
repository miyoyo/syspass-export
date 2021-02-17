package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s url apikey password", os.Args[0])
		os.Exit(1)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	search, err := MakeJSONRPCSearch(os.Args[2]).Marshal()
	if err != nil {
		panic("Could not create search query: " + err.Error())
	}

	response, err := http.Post(os.Args[1], "application/json", bytes.NewReader(search))

	if err != nil {
		panic("Could not talk to sysPass: " + err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic("Unable to parse body: " + err.Error())
	}

	results, err := UnmarshalAccountSearch(bodyBytes)

	if err != nil {
		panic("Unable to parse account search response: " + err.Error())
	}

	fmt.Printf("Found %d accounts:\n", len(results.Result.Result))

	file, _ := os.OpenFile("./db.kdbx", os.O_RDWR|os.O_CREATE, os.ModeAppend)
	defer file.Close()

	kdb := gokeepasslib.NewDatabase()
	kdb.Credentials = gokeepasslib.NewPasswordCredentials("TestPassword")

	rootGroup := gokeepasslib.NewGroup()
	rootGroup.Name = "sysPass"
	rootGroup.IconID = 48

	for i, s := range results.Result.Result {
		var group *gokeepasslib.Group

		for _, j := range rootGroup.Groups {
			if j.Name == s.CategoryName {
				group = &j
			}
		}

		if group == nil {
			temp := gokeepasslib.NewGroup()
			temp.Name = s.CategoryName
			rootGroup.Groups = append(rootGroup.Groups, temp)
			group = &temp
		}

		pass, err := GetPasswordForAccount(os.Args[1], os.Args[2], os.Args[3], s.ID)

		if err != nil {
			panic("Could not get password for account: " + err.Error())
		}

		kdb.Content.Meta.DatabaseName = "Database exported from sysPass at " + os.Args[1]

		fmt.Printf("%d: %s (%s), in group %s, for client %s with password %s\n", i, s.Name, s.Login, s.CategoryName, s.ClientName, pass)

		group.Entries = append(group.Entries, Entry("("+s.ClientName+") "+s.Name, s.Login, pass, s.URL))

		rootGroup.Groups = append(rootGroup.Groups, *group)
	}

	kdb.Content.Root.Groups = []gokeepasslib.Group{rootGroup}

	err = gokeepasslib.NewEncoder(file).Encode(kdb)

	if err != nil {
		panic("Failed to write keepass database: " + err.Error())
	}
}

func GetPasswordForAccount(url string, apiKey string, password string, id int64) (string, error) {
	reqBody, err := MakeJSONRPCViewPass(apiKey, password, id).Marshal()

	if err != nil {
		return "", fmt.Errorf("Unable to generate a password view request: ", err)
	}

	response, err := http.Post(url, "application/json", bytes.NewReader(reqBody))

	if err != nil {
		return "", fmt.Errorf("Could not contact syspass: ", err)
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("Unable to parse body: ", err.Error())
	}

	fmt.Println("%s\n", string(bodyBytes))

	pass, err := UnmarshalPasswordSearch(bodyBytes)

	if err != nil {
		return "", fmt.Errorf("Unable to parse password search response: ", err)
	}

	return pass.Result.Result.Password, nil

}

func Entry(name string, login string, password string, url string) gokeepasslib.Entry {
	entry := gokeepasslib.NewEntry()
	entry.Values = append(entry.Values, gokeepasslib.ValueData{
		Key: "Title",
		Value: gokeepasslib.V{
			Content:   name,
			Protected: wrappers.NewBoolWrapper(false),
		},
	})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{
		Key: "UserName",
		Value: gokeepasslib.V{
			Content:   login,
			Protected: wrappers.NewBoolWrapper(false),
		},
	})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{
		Key: "Password",
		Value: gokeepasslib.V{
			Content:   password,
			Protected: wrappers.NewBoolWrapper(false),
		},
	})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{
		Key: "URL",
		Value: gokeepasslib.V{
			Content:   url,
			Protected: wrappers.NewBoolWrapper(false),
		},
	})
	return entry
}
