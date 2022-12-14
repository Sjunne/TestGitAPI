package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Githubdata struct {
	Action string `json:"action"`
	Issue  struct {
		URL           string `json:"url"`
		RepositoryURL string `json:"repository_url"`
		LabelsURL     string `json:"labels_url"`
		CommentsURL   string `json:"comments_url"`
		EventsURL     string `json:"events_url"`
		HTMLURL       string `json:"html_url"`
		ID            int    `json:"id"`
		NodeID        string `json:"node_id"`
		Number        int    `json:"number"`
		Title         string `json:"title"`
		User          struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Labels            []interface{} `json:"labels"`
		State             string        `json:"state"`
		Locked            bool          `json:"locked"`
		Assignee          interface{}   `json:"assignee"`
		Assignees         []interface{} `json:"assignees"`
		Milestone         interface{}   `json:"milestone"`
		Comments          int           `json:"comments"`
		CreatedAt         time.Time     `json:"created_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		ClosedAt          interface{}   `json:"closed_at"`
		AuthorAssociation string        `json:"author_association"`
		ActiveLockReason  interface{}   `json:"active_lock_reason"`
		Draft             bool          `json:"draft"`
		PullRequest       struct {
			URL      string      `json:"url"`
			HTMLURL  string      `json:"html_url"`
			DiffURL  string      `json:"diff_url"`
			PatchURL string      `json:"patch_url"`
			MergedAt interface{} `json:"merged_at"`
		} `json:"pull_request"`
		Body      interface{} `json:"body"`
		Reactions struct {
			URL        string `json:"url"`
			TotalCount int    `json:"total_count"`
			Num1       int    `json:"+1"`
			Num10      int    `json:"-1"`
			Laugh      int    `json:"laugh"`
			Hooray     int    `json:"hooray"`
			Confused   int    `json:"confused"`
			Heart      int    `json:"heart"`
			Rocket     int    `json:"rocket"`
			Eyes       int    `json:"eyes"`
		} `json:"reactions"`
		TimelineURL           string      `json:"timeline_url"`
		PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
		StateReason           interface{} `json:"state_reason"`
	} `json:"issue"`
	Comment struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		IssueURL string `json:"issue_url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		User     struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		AuthorAssociation string    `json:"author_association"`
		Body              string    `json:"body"`
		Reactions         struct {
			URL        string `json:"url"`
			TotalCount int    `json:"total_count"`
			Num1       int    `json:"+1"`
			Num10      int    `json:"-1"`
			Laugh      int    `json:"laugh"`
			Hooray     int    `json:"hooray"`
			Confused   int    `json:"confused"`
			Heart      int    `json:"heart"`
			Rocket     int    `json:"rocket"`
			Eyes       int    `json:"eyes"`
		} `json:"reactions"`
		PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
	} `json:"comment"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL                  string        `json:"html_url"`
		Description              interface{}   `json:"description"`
		Fork                     bool          `json:"fork"`
		URL                      string        `json:"url"`
		ForksURL                 string        `json:"forks_url"`
		KeysURL                  string        `json:"keys_url"`
		CollaboratorsURL         string        `json:"collaborators_url"`
		TeamsURL                 string        `json:"teams_url"`
		HooksURL                 string        `json:"hooks_url"`
		IssueEventsURL           string        `json:"issue_events_url"`
		EventsURL                string        `json:"events_url"`
		AssigneesURL             string        `json:"assignees_url"`
		BranchesURL              string        `json:"branches_url"`
		TagsURL                  string        `json:"tags_url"`
		BlobsURL                 string        `json:"blobs_url"`
		GitTagsURL               string        `json:"git_tags_url"`
		GitRefsURL               string        `json:"git_refs_url"`
		TreesURL                 string        `json:"trees_url"`
		StatusesURL              string        `json:"statuses_url"`
		LanguagesURL             string        `json:"languages_url"`
		StargazersURL            string        `json:"stargazers_url"`
		ContributorsURL          string        `json:"contributors_url"`
		SubscribersURL           string        `json:"subscribers_url"`
		SubscriptionURL          string        `json:"subscription_url"`
		CommitsURL               string        `json:"commits_url"`
		GitCommitsURL            string        `json:"git_commits_url"`
		CommentsURL              string        `json:"comments_url"`
		IssueCommentURL          string        `json:"issue_comment_url"`
		ContentsURL              string        `json:"contents_url"`
		CompareURL               string        `json:"compare_url"`
		MergesURL                string        `json:"merges_url"`
		ArchiveURL               string        `json:"archive_url"`
		DownloadsURL             string        `json:"downloads_url"`
		IssuesURL                string        `json:"issues_url"`
		PullsURL                 string        `json:"pulls_url"`
		MilestonesURL            string        `json:"milestones_url"`
		NotificationsURL         string        `json:"notifications_url"`
		LabelsURL                string        `json:"labels_url"`
		ReleasesURL              string        `json:"releases_url"`
		DeploymentsURL           string        `json:"deployments_url"`
		CreatedAt                time.Time     `json:"created_at"`
		UpdatedAt                time.Time     `json:"updated_at"`
		PushedAt                 time.Time     `json:"pushed_at"`
		GitURL                   string        `json:"git_url"`
		SSHURL                   string        `json:"ssh_url"`
		CloneURL                 string        `json:"clone_url"`
		SvnURL                   string        `json:"svn_url"`
		Homepage                 interface{}   `json:"homepage"`
		Size                     int           `json:"size"`
		StargazersCount          int           `json:"stargazers_count"`
		WatchersCount            int           `json:"watchers_count"`
		Language                 string        `json:"language"`
		HasIssues                bool          `json:"has_issues"`
		HasProjects              bool          `json:"has_projects"`
		HasDownloads             bool          `json:"has_downloads"`
		HasWiki                  bool          `json:"has_wiki"`
		HasPages                 bool          `json:"has_pages"`
		ForksCount               int           `json:"forks_count"`
		MirrorURL                interface{}   `json:"mirror_url"`
		Archived                 bool          `json:"archived"`
		Disabled                 bool          `json:"disabled"`
		OpenIssuesCount          int           `json:"open_issues_count"`
		License                  interface{}   `json:"license"`
		AllowForking             bool          `json:"allow_forking"`
		IsTemplate               bool          `json:"is_template"`
		WebCommitSignoffRequired bool          `json:"web_commit_signoff_required"`
		Topics                   []interface{} `json:"topics"`
		Visibility               string        `json:"visibility"`
		Forks                    int           `json:"forks"`
		OpenIssues               int           `json:"open_issues"`
		Watchers                 int           `json:"watchers"`
		DefaultBranch            string        `json:"default_branch"`
	} `json:"repository"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)
	test, _ := io.ReadAll(r.Body)

	myObj := Githubdata{}
	err := json.Unmarshal(test, &myObj)
	if err != nil {
		return
	}
	fmt.Printf("headers: %v\n", myObj)
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//JSON vi kan bruge:
//login: brugernavn
//X-Github-Event:[issue_comment]
//"node_id":"MDQ6VXNlcjg5NTcyODg="
//"name":"TestGitAPI"
//"full_name":"Sjunne/TestGitAPI"

