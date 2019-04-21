package azuredevops

import (
	"fmt"
	"net/url"
	"time"
)

// VersionControlChangeType enum declaration
type VersionControlChangeType int

// VersionControlChangeType valid enum values
const (
	None VersionControlChangeType = iota
	Add
	Edit
	Encoding
	Rename
	Delete
	Undelete
	Branch
	Merge
	Lock
	Rollback
	SourceRename
	TargetRename
	Property
	All
)

func (d VersionControlChangeType) String() string {
	return [...]string{"none", "add", "edit", "encoding", "rename", "delete", "undelete", "branch", "merge", "lock", "rollback", "sourceRename", "targetRename", "property", "all"}[d]
}

// GitObjectType enum declaration
type GitObjectType int

// GitObjectType enum declaration
const (
	Bad GitObjectType = iota
	Commit
	Tree
	Blob
	Tag
	Ext2
	OfsDelta
	RefDelta
)

func (d GitObjectType) String() string {
	return [...]string{"bad", "commit", "tree", "blob", "tag", "ext2", "ofsDelta", "refDelta"}[d]
}

// GitService handles communication with the git methods on the API
// See: https://docs.microsoft.com/en-us/rest/api/vsts/git/
type GitService struct {
	client *Client
}

// FileContentMetadata Describes files referenced by a GitItem
type FileContentMetadata struct {
	ContentType *string `json:"contentType,omitempty"`
	Encoding    *int    `json:"encoding,omitempty"`
	Extension   *string `json:"extension,omitempty"`
	FileName    *string `json:"fileName,omitempty"`
	IsBinary    *bool   `json:"isBinary,omitempty"`
	IsImage     *bool   `json:"isImage,omitempty"`
	VSLink      *string `json:"vsLink,omitempty"`
}

// GitRefsResponse describes the git list refs response
type GitRefsResponse struct {
	Count   int       `json:"count"`
	GitRefs []*GitRef `json:"value"`
}

// GitStatusesResponse describes the git statuses response
type GitStatusesResponse struct {
	Count       int          `json:"count"`
	GitStatuses []*GitStatus `json:"value"`
}

// GitChange describes file path and content changes
type GitChange struct {
	ChangeID           *int                      `json:"changeId,omitempty"`
	ChangeType         *VersionControlChangeType `json:"changeType,omitempty"`
	Item               *GitItem                  `json:"item,omitempty"`
	NewContent         *ItemContent              `json:"newContent,omitempty"`
	NewContentTemplate *GitTemplate              `json:"newContentTemplate,omitempty"`
	OriginalPath       *string                   `json:"originalPath,omitempty"`
	SourceServerItem   *string                   `json:"sourceServerItem,omitempty"`
	URL                *string                   `json:"url,omitempty"`
}

// GitCommitChanges is a list of GitCommitRefs and count of all changes describes in
// the response from the API
type GitCommitChanges struct {
	ChangeCounts *ChangeCountDictionary `json:"changeCounts,omitempty"`
	Changes      *[]GitChange           `json:"changes,omitempty"`
}

// GitCommitRef describes a single git commit reference
type GitCommitRef struct {
	Links            *[]ReferenceLinks      `json:"_links,omitempty"`
	CommitID         *string                `json:"commitId,omitempty"`
	Author           *GitUserDate           `json:"author,omitempty"`
	Committer        *GitUserDate           `json:"committer,omitempty"`
	Comment          *string                `json:"comment,omitempty"`
	CommentTruncated *bool                  `json:"commentTruncated,omitempty"`
	URL              *string                `json:"url,omitempty"`
	ChangeCounts     *ChangeCountDictionary `json:"changeCounts,omitempty"`
	Changes          *GitChange             `json:"changes,omitempty"`
	Parents          *[]string              `json:"parents,omitempty"`
	Push             *GitPushRef            `json:"push,omitempty"`
	RemoteURL        *string                `json:"remoteUrl,omitempty"`
	Statuses         *[]GitStatus           `json:"statuses,omitempty"`
	WorkItems        *ResourceRef           `json:"workItems,omitempty"`
}

// GitRef provides information about a git/fork ref.
type GitRef struct {
	Links          *[]ReferenceLinks `json:"_links,omitempty"`
	Creator        *IdentityRef      `json:"creator,omitempty"`
	IsLocked       *bool             `json:"isLocked,omitempty"`
	IsLockedBy     *IdentityRef      `json:"isLockedBy,omitempty"`
	Name           *string           `json:"name,omitempty"`
	ObjectID       *string           `json:"objectId,omitempty"`
	PeeledObjectID *string           `json:"peeledObjectId,omitempty"`
	Repository     *GitRepository    `json:"repository,omitempty"`
	Statuses       *[]GitStatus      `json:"statuses,omitempty"`
	URL            *string           `json:"url,omitempty"`
}

// GitItem describes a single git item
type GitItem struct {
	Links                 *[]ReferenceLinks    `json:"_links,omitempty"`
	CommitID              *string              `json:"commitId,omitempty"`
	Content               *string              `json:"content,omitempty"`
	ContentMetadata       *FileContentMetadata `json:"contentMetadata,omitempty"`
	GitObjectType         *GitObjectType       `json:"gitObjectType,omitempty"`
	IsFolder              *bool                `json:"isFolder,omitempty"`
	IsSymLink             *bool                `json:"isSymLink,omitempty"`
	LatestProcessedChange *GitCommitRef        `json:"latestProcessedChange,omitempty"`
	ObjectID              *string              `json:"objectId,omitempty"`
	OriginalObjectID      *string              `json:"originalObjectId,omitempty"`
	Path                  *string              `json:"path,omitempty"`
	URL                   *string              `json:"url,omitempty"`
}

// GitPullRequest represents all the data associated with a pull request.
type GitPullRequest struct {
	Links                 *[]ReferenceLinks                `json:"_links,omitempty"`
	ArtifactID            *string                          `json:"artifactId,omitempty"`
	AutoCompleteSetBy     *IdentityRef                     `json:"autoCompleteSetBy,omitempty"`
	ClosedBy              *IdentityRef                     `json:"closedBy,omitempty"`
	ClosedDate            *time.Time                       `json:"closedDate,omitempty"`
	CodeReviewID          *int                             `json:"codeReviewId,omitempty"`
	Commits               *[]GitCommitRef                  `json:"commits,omitempty"`
	CompletionOptions     *GitPullRequestCompletionOptions `json:"completionOptions,omitempty"`
	CompletionQueueTime   *time.Time                       `json:"completionQueueTime,	omitempty"`
	CreatedBy             *IdentityRef                     `json:"createdBy,omitempty"`
	CreationDate          *time.Time                       `json:"creationDate,omitempty"`
	Description           *string                          `json:"description,omitempty"`
	ForkSource            *GitRef                          `json:"forkSource,omitempty"`
	IsDraft               *bool                            `json:"isDraft,omitempty"`
	Labels                *WebAPITagDefinition             `json:"labels,omitempty"`
	LastMergeCommit       *GitCommitRef                    `json:"lastMergeCommit,omitempty"`
	LastMergeSourceCommit *GitCommitRef                    `json:"lastMergeSourceCommit,omitempty"`
	LastMergeTargetCommit *GitCommitRef                    `json:"lastMergeTargetCommit,omitempty"`
	MergeFailureMessage   *string                          `json:"mergeFailureMessage,omitempty"`
	MergeFailureType      *PullRequestMergeFailureType     `json:"mergeFailureType,omitempty"`
	MergeID               *string                          `json:"mergeId,omitempty"`
	MergeOptions          *GitPullRequestMergeOptions      `json:"mergeOptions,omitempty"`
	MergeStatus           *PullRequestAsyncStatus          `json:"mergeStatus,omitempty"`
	PullRequestID         *int                             `json:"pullRequestId,omitempty"`
	Repository            *GitRepository                   `json:"repository,omitempty"`
	Reviewers             *[]IdentityRefWithVote           `json:"reviewers,omitempty"`
	RemoteURL             *string                          `json:"remoteUrl,omitempty"`
	SourceRefName         *string                          `json:"sourceRefName,omitempty"`
	Status                *PullRequestStatus               `json:"status,omitempty"`
	SupportsIterations    *bool                            `json:"supportsIterations,omitempty"`
	TargetRefName         *string                          `json:"targetRefName,omitempty"`
	Title                 *string                          `json:"title,omitempty"`
	URL                   *string                          `json:"url,omitempty"`
	WorkItemRefs          *[]ResourceRef                   `json:"workItemRefs,omitempty"`
}

// GitPullRequestCompletionOptions describes preferences about how the pull
// request should be completed.
// SquashMerge is deprecated. You should explicity set the value of MergeStrategy. If
// MergeStrategy is set to any value, the SquashMerge value will be ignored. If
// MergeStrategy is not set, the merge strategy will be no-fast-forward if this flag is false, or squash if true.
// https://docs.microsoft.com/en-us/rest/api/azure/devops/git/pull%20requests/update?view=azure-devops-rest-5.1#pullrequeststatus
type GitPullRequestCompletionOptions struct {
	BypassPolicy            bool                        `json:"bypassPolicy,omitempty"`
	BypassReason            string                      `json:"bypassReason,omitempty"`
	DeleteSourceBranch      bool                        `json:"deleteSourceBranch,omitempty"`
	MergeCommitMessage      string                      `json:"mergeCommitMessage,omitempty"`
	MergeStrategy           GitPullRequestMergeStrategy `json:"mergeStrategy,omitempty"`
	SquashMerge             bool                        `json:"squashMerge,omitempty"`
	TransitionWorkItems     bool                        `json:"transitionWorkItems,omitempty"`
	TriggeredByAutoComplete bool                        `json:"triggeredByAutoComplete,omitempty"`
}

// GitPullRequestMergeOptions describes the options which are used when a pull
// request merge is created.
type GitPullRequestMergeOptions struct {
	DetectRenameFalsePositives *bool `json:"detectRenameFalsePositives,omitempty"`
	DisableRenames             *bool `json:"disableRenames,omitempty"`
}

// GitPullRequestMergeStrategy specifies the strategy used to merge the pull request
// during completion.
type GitPullRequestMergeStrategy int

// GitPullRequestMergeStrategy enum values
const (
	NoFastForward GitPullRequestMergeStrategy = iota
	Rebase
	SebaseMerge
	Squash
)

func (d GitPullRequestMergeStrategy) String() string {
	return [...]string{"noFastForward", "rebase", "rebaseMerge", "squash"}[d]
}

// GitPush describes a code push request event.
type GitPush struct {
	Links      *[]ReferenceLinks `json:"_links,omitempty"`
	Commits    *[]GitCommitRef   `json:"commits,omitempty"`
	Date       *time.Time        `json:"date,omitempty"`
	PushID     *int              `json:"pushId,omitempty"`
	PushedBy   *IdentityRef      `json:"pushedBy,omitempty"`
	RefUpdates *[]GitRefUpdate   `json:"refUpdates,omitempty"`
	Repository *GitRepository    `json:"repository,omitempty"`
	URL        *string           `json:"url,omitempty"`
}

// GitPushRef Describes a push request
type GitPushRef struct {
	Commits    *[]GitCommitRef `json:"commits,omitempty"`
	RefUpdates *[]GitRefUpdate `json:"refUpdates,omitempty"`
	Repository *GitRepository  `json:"repository,omitempty"`
}

// GitRefUpdate
type GitRefUpdate struct {
	IsLocked     *bool   `json:"isLocked,omitempty"`
	Name         *string `json:"name,omitempty"`
	NewObjectID  *string `json:"newObjectId,omitempty"`
	OldObjectID  *string `json:"oldObjectId,omitempty"`
	RepositoryID *string `json:"repositoryId,omitempty"`
}

// GitRepository describes an Azure Devops Git repository.
type GitRepository struct {
	Links            *[]ReferenceLinks     `json:"_links,omitempty"`
	DefaultBranch    *string               `json:"defaultBranch,omitempty"`
	ID               *string               `json:"id,omitempty"`
	IsFork           *bool                 `json:"isFork,omitempty"`
	Name             *string               `json:"name,omitempty"`
	ParentRepository *GitRepositoryRef     `json:"parentRepository,omitempty"`
	Project          *TeamProjectReference `json:"project,omitempty"`
	RemoteURL        *string               `json:"remoteUrl,omitempty"`
	Size             *int                  `json:"size,omitempty"`
	SSHURL           *string               `json:"sshUrl,omitempty"`
	URL              *string               `json:"url,omitempty"`
	ValidRemoteURLs  *[]string             `json:"validRemoteUrls,omitempty"`
}

// GitRepositoryRef
type GitRepositoryRef struct {
	Collection *TeamProjectCollectionReference `json:"collection,omitempty"`
	ID         *string                         `json:"id,omitempty"`
	IsFork     *bool                           `json:"isFork,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Project    *TeamProjectReference           `json:"project,omitempty"`
	RemoteURL  *string                         `json:"remoteUrl,omitempty"`
	SSHURL     *string                         `json:"sshUrl,omitempty"`
	URL        *string                         `json:"url,omitempty"`
}

// GitStatusState contains the metadata of a service/extension posting a status.
type GitStatusState int

// GitStatusState enum values
const (
	GitNotSet GitStatusState = iota
	GitPending
	GitSucceeded
	GitFailed
	GitError
	GitNotApplicable
)

func (d GitStatusState) String() string {
	return [...]string{"notSet", "pending", "succeeded", "failed", "error", "notApplicable"}[d]
}

type GitStatus struct {
	Links        *[]ReferenceLinks `json:"_links,omitempty"`
	Context      *GitStatusContext `json:"context,omitempty"`
	CreatedBy    *IdentityRef      `json:"createdBy,omitempty"`
	CreationDate *time.Time        `json:"creationDate,omitempty"`
	Description  *string           `json:"description,omitempty"`
	ID           *int              `json:"id,omitempty"`
	State        *GitStatusState   `json:"state,omitempty"`
	TargetURL    *string           `json:"targetUrl,omitempty"`
	UpdatedDate  *time.Time        `json:"updatedDate,omitempty"`
}

// GitPullRequestStatus This class contains the metadata of a service/extension
// posting pull request status. Status can be associated with a pull request or
// an iteration.
type GitPullRequestStatus struct {
	GitStatus
	IterationID int        `json:"iterationId,omitempty"`
	Properties  *time.Time `json:"properties,omitempty"`
}

// GitRefListOptions describes what the request to the API should look like
type GitRefListOptions struct {
	Filter             string `url:"filter,omitempty"`
	IncludeStatuses    bool   `url:"includeStatuses,omitempty"`
	LatestStatusesOnly bool   `url:"latestStatusesOnly,omitempty"`
}

// GitStatusContext Status context that uniquely identifies the status.
type GitStatusContext struct {
	Genre *string `json:"genre,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// GitTemplate
type GitTemplate struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// GitUserDate User info and date for Git operations.
type GitUserDate struct {
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
	Date  *time.Time `json:"date,omitempty"`
}

// ListRefs returns a list of the references for a git repo
func (s *GitService) ListRefs(repo, refType string, opts *GitRefListOptions) ([]*GitRef, int, error) {
	URL := fmt.Sprintf(
		"_apis/git/repositories/%s/refs/%s?api-version=%s",
		repo,
		refType,
		APIVersion,
	)

	URL, err := addOptions(URL, opts)

	request, err := s.client.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, 0, err
	}
	var response GitRefsResponse
	_, err = s.client.Execute(request, &response)

	return response.GitRefs, response.Count, err
}

// GetRepository Return a single GitRepository
// https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories/get%20repository?view=azure-devops-rest-5.1
func (s *GitService) GetRepository(repoName string) (*GitRepository, int, error) {
	URL := fmt.Sprintf(
		"_apis/git/repositories/%s?api-version=%s",
		repoName,
		APIVersion,
	)

	request, err := s.client.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, 0, err
	}
	var response GitRepository
	_, err = s.client.Execute(request, &response)

	return &response, 1, err

}

// CreateStatus creates a new status for a repository at the specified
// reference. Ref can be a SHA, a branch name, or a tag name.
// https://docs.microsoft.com/en-us/rest/api/azure/devops/git/statuses/create?view=azure-devops-rest-5.0
func (s *GitService) CreateStatus(repoName, ref string, status GitStatus) (*GitStatus, int, error) {
	URL := fmt.Sprintf(
		"_apis/git/repositories/%s/commits/%s/statuses?api-version=%s",
		url.QueryEscape(ref),
		ref,
		APIVersion,
	)

	request, err := s.client.NewRequest("POST", URL, nil)
	if err != nil {
		return nil, 0, err
	}
	var response GitStatus
	_, err = s.client.Execute(request, &response)

	return &response, 1, err
}
