package api

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"log"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	//"github.com/go-git/go-git/v5/plumbing"
)

func GithubInfoDeep(username string, fork bool) EmailsType {
	log.Println("github")
	var data []struct {
		//Id     string `json:"id"`
		//NodeId string `json:"node_id"`
		Name string `json:"name"`

		FullName   string `json:"full_name"`
		Fork       bool   `json:"fork"`
		Url        string `json:"url"`
		GitUrl     string `json:"git_url"`
		SshUrl     string `json:"ssh_url"`
		CloneUrl   string `json:"clone_url"`
		OpenIssues int    `json:"open_issues"`
		Forks      int    `json:"forks"`
		Homepage   string `json:"homepage"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
		Pushed_at  string `json:"pushed_at"`
	}

	fatal := false
	jsonData, err := HttpRequest("https://api.github.com/users/" + username + "/repos")
	if err != nil {
		log.Println(jsonData)
		fatal = true
		return EmailsType{"fatal": Email{}}
	}

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
		log.Println("propably rate limited")
		fatal = true
		return EmailsType{"fatal": Email{}}
	} else {

		contributors := make(map[string]bool)
		foundEmail := make(map[string]bool)
		for _, repo := range data {
			//if repo.Fork == fork || repo.Fork {
			log.Println(repo.Name)

			r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
				URL: repo.CloneUrl,
			})
			Check(err)
			//head, err := r.Head()
			//Check(err)
			//commitIter, err := r.Log(&git.LogOptions{From: head.Hash()})

			commitIter, err := r.Log(&git.LogOptions{})
			Check(err)
			err = commitIter.ForEach(func(c *object.Commit) error {
				if !contributors[c.Author.Email] && !IsGitHubMail(c.Author.Email) {
					type Author struct {
						Name  string `json:"name"`
						Email string `json:"email"`
					}
					var commitInfo struct {
						Author Author `json:"author"`
					}

					jsonData, err := HttpRequest(fmt.Sprintf("https://api.github.com/repos/%s/git/commits/%s", repo.FullName, c.Hash.String()))
					if err != nil {
						fatal = true
					} else {

						err = json.Unmarshal([]byte(jsonData), &commitInfo)
						if err != nil {
							log.Println(err)
							fatal = true
						}
						log.Printf("Author: %s\nUsername: %s\n", commitInfo.Author.Name, username)
						if strings.EqualFold(commitInfo.Author.Name, username) { // check username
							log.Println("found:")
							log.Println(c.Author.Email)
							foundEmail[c.Author.Email] = true
						}
					}
					contributors[c.Author.Email] = true
				}
				//log.Println(c.Hash.String())
				return nil
			})
			Check(err)

		}
		if fatal {
			return EmailsType{}
		}
		foundEmailArray := EmailsType{}
		for c := range foundEmail {
			foundEmailArray[c] = Email{
				Mail: c,
				Src:  "github",
				Services: EmailServiceEnums{
					"github": {
						Name:     "GitHub",
						Username: username,
						Link:     fmt.Sprintf("https://github.com/%s", username),
						Icon:     "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
					},
				},
			}
		}
		//}
		return foundEmailArray
	}
	return EmailsType{}
}
