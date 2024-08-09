package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v62/github"
)

type Github struct {
	GithubAppID    string
	GithubClientID string
	GithubSecret   string
}

type UploadStruct struct {
	Repo    string `json:"repo"`
	Content string `json:"content"`
	File    string `json:"file"`
	Owner   string `json:"owner"`
}

type GitHubMeStruct struct {
	Login string `json:"login"`
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

type Repo struct {
	Name  string
	Owner string
}

type Dir struct {
	Path string
	Type string
}

// func GetGithubEnv() Github {
// 	return Github{
// 		os.Getenv("GITHUB_APP_ID"),
// 		os.Getenv("GITHUB_CLIENT_ID"),
// 		os.Getenv("GITHUB_SECRET"),
// 	}
// }

// func GithubRedirect(c *fiber.Ctx) error {
// 	// create a CSRF token and store it locally
// 	b := make([]byte, 16)
// 	_, _ = rand.Read(b)
// 	state := hex.EncodeToString(b)
// 	c.Cookie(&fiber.Cookie{
// 		Name:  "github_state",
// 		Value: state,
// 	})
// 	scope := "repo:status+repo+public_repo+admin:repo_hook+admin:org+admin:public_key+admin:org_hook+user+user:follow+read:gpg_key+delete:packages+read:discussion+workflow+admin:gpg_key+write:packages+delete_repo+read:user+gist+write:public_key+write:org+write:repo_hook+repo:invite+write:gpg_key+read:packages+write:discussion+user:email+notifications+read:public_key+read:org+read:repo_hook+security_events+repo_deployment"

// 	d := GetGithubEnv()
// 	link := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&response_type=code&scope=%s&redirect_uri=%s&state=%s",
// 		d.GithubClientID,
// 		scope,
// 		fmt.Sprintf("%s/callback", os.Getenv("DOMAIN_NAME")),
// 		state,
// 	)
// 	return c.Redirect(link, http.StatusTemporaryRedirect)
// }

// func GithubCallback(c *fiber.Ctx) error {
// 	code := c.Query("code")
// 	if c.Query("state") != c.Cookies("github_state") {
// 		return c.Status(400).SendString("error - bad state")
// 	}

//		tkn := getGithubAccessToken(code)
//		if tkn == "" {
//			return c.Status(400).SendString("error")
//		}
//		c.Cookie(&fiber.Cookie{
//			Name:  "tkn",
//			Value: tkn,
//		})
//		return c.Redirect("/repos", http.StatusTemporaryRedirect)
//	}
func isMarkdownFile(filename string) bool {
	ext := filepath.Ext(filename)
	return strings.ToLower(ext) == ".md"
}

func GithubSetToken(c *fiber.Ctx) error {
	tkn := c.Query("tkn")
	data := getGithubData(tkn)
	if data.Id == 0 || data.Login == "" {
		c.Status(400).JSON(fiber.Map{"error": "bad token"})
		return nil
	}
	c.Cookie(&fiber.Cookie{
		Name:  "tkn",
		Value: tkn,
	})
	c.Status(200).JSON(fiber.Map{"status": "success"})
	return nil
}

func GithubMyRepos(c *fiber.Ctx) error {
	tkn := c.Cookies("tkn")
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	client := github.NewClient(nil).WithAuthToken(tkn)
	listOpts := github.ListOptions{Page: pageInt, PerPage: 10}
	opts := github.RepositoryListByAuthenticatedUserOptions{Sort: "updated", Affiliation: "owner,collaborator,organization_member", ListOptions: listOpts}
	list, _, err := client.Repositories.ListByAuthenticatedUser(context.Background(), &opts)
	if err != nil {
		return c.Status(200).SendString(fmt.Sprintf("error, %e", err))
	}

	r := []Repo{}
	for _, l := range list {
		r = append(r, Repo{Name: *l.Name, Owner: *l.Owner.Login})
	}
	return c.Status(200).JSON(r)
}

func GithubRepoFiles(c *fiber.Ctx) error {
	tkn := c.Cookies("tkn")
	repoName := c.Query("repo")
	owner := c.Query("owner")

	repoParts := strings.SplitN(repoName, "/", 2)
	var path string
	if len(repoParts) > 1 {
		path = "/" + repoParts[1]
	} else {
		path = "/"
	}
	res, err := getRepoFiles(tkn, repoParts[0], owner, path)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error, %e", err))
	}
	return c.Status(200).JSON(res)
}

func GithubRepoFile(c *fiber.Ctx) error {
	tkn := c.Cookies("tkn")
	owner := c.Query("owner")
	repoName := c.Query("repo")
	path := c.Query("path")

	decodedPath, err := url.QueryUnescape(path)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return err
	}

	res, err := getFile(tkn, owner, repoName, decodedPath)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error, %e", err))
	}
	return c.Status(200).SendString(res)
}

func GithubSendFile(c *fiber.Ctx) error {
	tkn := c.Cookies("tkn")
	payload := UploadStruct{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	decodedPath, err := url.QueryUnescape(payload.File)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return err
	}
	err = sendFile(tkn, payload.Owner, payload.Repo, payload.Content, decodedPath, fmt.Sprintf("Update %s", decodedPath))
	log.Println(err)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error, %e", err))
	}
	return c.Status(200).SendString("success")
}

func getRepoFiles(tkn, repo, owner, path string) (res []Dir, err error) {
	// data := getGithubData(tkn)
	client := github.NewClient(nil).WithAuthToken(tkn)
	ctx := context.Background()
	opts := github.RepositoryContentGetOptions{}
	_, d, resp, err := client.Repositories.GetContents(ctx, owner, repo, path, &opts)

	if resp.StatusCode != 200 || err != nil {
		return res, errors.New("File does not exist")
	}
	data := []Dir{}
	for _, fd := range d {
		if fd.GetType() == "file" {
			if isMarkdownFile(fd.GetName()) {
				data = append(data, Dir{Path: fd.GetPath(), Type: fd.GetType()})
			}
		} else if fd.GetType() == "dir" {
			data = append(data, Dir{Path: fd.GetPath(), Type: fd.GetType()})
		}
	}
	return data, nil
}

func getFile(tkn, owner, repo, fileName string) (string, error) {
	client := github.NewClient(nil).WithAuthToken(tkn)
	ctx := context.Background()

	f, _, resp, err := client.Repositories.GetContents(ctx, owner, repo, fileName, nil)
	if resp.StatusCode != 200 || err != nil || f.Content == nil {
		return "", errors.New("File does not exist")
	}
	res, err := f.GetContent()
	if err != nil {
		return "", err
	}
	return res, nil
}

func sendFile(tkn, owner, repo, content, Path, commitMessage string) error {
	data := getGithubData(tkn)
	client := github.NewClient(nil).WithAuthToken(tkn)
	ctx := context.Background()
	f, _, resp, err := client.Repositories.GetContents(ctx, owner, repo, Path, nil)
	if resp.StatusCode != 200 || err != nil {
		opts := &github.RepositoryContentFileOptions{
			Message:   github.String(commitMessage),
			Content:   []byte(content),
			Committer: &github.CommitAuthor{Name: github.String(data.Login), Email: github.String(data.Email)},
		}
		repoParts := strings.SplitN(Path, "/", 2)
		commit, resp, err := client.Repositories.CreateFile(ctx, owner, repo, repoParts[1], opts)
		if resp.StatusCode != 201 {
			return errors.New(fmt.Sprintf("Error updating file: %s", resp.Status))
		} else if err != nil {
			return err
		}
		fmt.Println(commit)
		return nil
	} else {
		opts := &github.RepositoryContentFileOptions{
			Message:   github.String(commitMessage),
			SHA:       f.SHA,
			Content:   []byte(content),
			Committer: &github.CommitAuthor{Name: github.String(data.Login), Email: github.String(data.Email)},
		}
		commit, resp, err := client.Repositories.UpdateFile(ctx, owner, repo, *f.Path, opts)
		if resp.StatusCode != 200 {
			return errors.New(fmt.Sprintf("Error updating file: %s", resp.Status))
		} else if err != nil {
			return err
		}
		fmt.Println(commit)
		return nil
	}
}

// func getGithubAccessToken(code string) string {
// 	d := GetGithubEnv()
// 	// Set us the request body as JSON
// 	requestBodyMap := map[string]string{
// 		"client_id":     d.GithubClientID,
// 		"client_secret": d.GithubSecret,
// 		"code":          code,
// 	}
// 	requestJSON, _ := json.Marshal(requestBodyMap)

// 	// POST request to set URL
// 	req, reqerr := http.NewRequest(
// 		"POST",
// 		"https://github.com/login/oauth/access_token",
// 		bytes.NewBuffer(requestJSON),
// 	)
// 	if reqerr != nil {
// 		log.Panic("Request creation failed")
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	// Get the response
// 	resp, resperr := http.DefaultClient.Do(req)
// 	if resperr != nil {
// 		log.Panic("Request failed")
// 	}

// 	// Response body converted to stringified JSON
// 	respbody, _ := ioutil.ReadAll(resp.Body)

// 	// Represents the response received from Github
// 	type githubAccessTokenResponse struct {
// 		AccessToken string `json:"access_token"`
// 		TokenType   string `json:"token_type"`
// 		Scope       string `json:"scope"`
// 	}

// 	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
// 	var ghresp githubAccessTokenResponse
// 	json.Unmarshal(respbody, &ghresp)
// 	fmt.Println("ghresp.Scope: ", ghresp.Scope)

// 	// Return the access token (as the rest of the
// 	// details are relatively unnecessary for us)
// 	return ghresp.AccessToken
// }

func getGithubData(accessToken string) GitHubMeStruct {
	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	// Set the Authorization header before sending the request
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Read the response as a byte slice
	respbody, _ := ioutil.ReadAll(resp.Body)
	res := GitHubMeStruct{}
	json.Unmarshal(respbody, &res)

	//return string(respbody)
	return res
}
