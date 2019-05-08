package jiracmd

import (
	"fmt"

	"github.com/coryb/figtree"
	"github.com/coryb/oreo"

	"gopkg.in/Netflix-Skunkworks/go-jira.v1"
	"github.com/albertrdixon/go-jira/jiracli"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1/jiradata"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type SubtaskOptions struct {
	jiracli.CommonOptions `yaml:",inline" json:",inline" figtree:",inline"`
	jiradata.IssueUpdate  `yaml:",inline" json:",inline" figtree:",inline"`
	Project               string            `yaml:"project,omitempty" json:"project,omitempty"`
	IssueType             string            `yaml:"issuetype,omitempty" json:"issuetype,omitempty"`
	Overrides             map[string]string `yaml:"overrides,omitempty" json:"overrides,omitempty"`
	Issue                 string            `yaml:"issue,omitempty" json:"issue,omitempty"`
}

func CmdSubtaskRegistry() *jiracli.CommandRegistryEntry {
	opts := SubtaskOptions{
		CommonOptions: jiracli.CommonOptions{
			Template: figtree.NewStringOption("subtask"),
		},
		Overrides: map[string]string{},
	}

	return &jiracli.CommandRegistryEntry{
		"Subtask issue",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdSubtaskUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			if opts.IssueType == "" {
				opts.IssueType = "Sub-task"
			}
			return CmdSubtask(o, globals, &opts)
		},
	}
}

func CmdSubtaskUsage(cmd *kingpin.CmdClause, opts *SubtaskOptions) error {
	jiracli.BrowseUsage(cmd, &opts.CommonOptions)
	jiracli.EditorUsage(cmd, &opts.CommonOptions)
	jiracli.TemplateUsage(cmd, &opts.CommonOptions)
	cmd.Flag("noedit", "Disable opening the editor").SetValue(&opts.SkipEditing)
	cmd.Flag("project", "project to subtask issue in").Short('p').StringVar(&opts.Project)
	cmd.Flag("comment", "Comment message for issue").Short('m').PreAction(func(ctx *kingpin.ParseContext) error {
		opts.Overrides["comment"] = jiracli.FlagValue(ctx, "comment")
		return nil
	}).String()
	cmd.Flag("override", "Set issue property").Short('o').StringMapVar(&opts.Overrides)
	cmd.Arg("ISSUE", "Parent issue for subtask").StringVar(&opts.Issue)
	return nil
}

// CmdSubtask sends the subtask-metadata to the "subtask" template for editing, then
// will parse the edited document as YAML and submit the document to jira.
func CmdSubtask(o *oreo.Client, globals *jiracli.GlobalOptions, opts *SubtaskOptions) error {
	type templateInput struct {
		Meta      *jiradata.IssueType `yaml:"meta" json:"meta"`
		Overrides map[string]string   `yaml:"overrides" json:"overrides"`
		Parent    *jiradata.Issue     `yaml:"parent" json:"parent"`
	}

	parent, err := jira.GetIssue(o, globals.Endpoint.Value, opts.Issue, nil)
	if err != nil {
		return err
	}

	if project, ok := parent.Fields["project"].(map[string]interface{}); ok {
		if key, ok := project["key"].(string); ok {
			opts.Project = key
		} else {
			return fmt.Errorf("Failed to find Project Key in parent issue")
		}
	} else {
		return fmt.Errorf("Failed to find Project field in parent issue")
	}

	createMeta, err := jira.GetIssueCreateMetaIssueType(o, globals.Endpoint.Value, opts.Project, opts.IssueType)
	if err != nil {
		return err
	}

	issueUpdate := jiradata.IssueUpdate{}
	input := templateInput{
		Meta:      createMeta,
		Overrides: opts.Overrides,
		Parent:    parent,
	}
	input.Overrides["project"] = opts.Project
	input.Overrides["issuetype"] = opts.IssueType
	input.Overrides["user"] = globals.User.Value

	var issueResp *jiradata.IssueCreateResponse
	err = jiracli.EditLoop(&opts.CommonOptions, &input, &issueUpdate, func() error {
		issueResp, err = jira.CreateIssue(o, globals.Endpoint.Value, &issueUpdate)
		return err
	})
	if err != nil {
		return err
	}

	if !globals.Quiet.Value {
		fmt.Printf("OK %s %s\n", issueResp.Key, jira.URLJoin(globals.Endpoint.Value, "browse", issueResp.Key))
	}

	if opts.Browse.Value {
		return CmdBrowse(globals, issueResp.Key)
	}
	return nil
}
