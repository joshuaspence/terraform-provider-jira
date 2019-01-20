package jira

import (
	"io/ioutil"
	"reflect"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

// API Endpoints
const groupAPIEndpoint = "/rest/api/2/group"
const groupUserAPIEndpoint = "/rest/api/2/group/user"

const issueLinkAPIEndpoint = "/rest/api/2/issueLink"
const issueLinkTypeAPIEndpoint = "/rest/api/2/issueLinkType"
const issueTypeAPIEndpoint = "/rest/api/2/issuetype"

const projectAPIEndpoint = "/rest/api/2/project"

func request(client *jira.Client, method string, endpoint string, in interface{}, out interface{}) error {

	req, err := client.NewRequest(method, endpoint, in)

	if err != nil {
		return errors.Wrapf(err, "Creating %s Request failed", method)
	}

	res, err := client.Do(req, out)
	if err != nil {
		typeName := reflect.TypeOf(in).Name()
		body, readErr := ioutil.ReadAll(res.Response.Body)
		if readErr != nil {
			return errors.Wrapf(readErr, "Creating %s Request failed", typeName)
		}
		return errors.Wrapf(err, "Creating %s Request failed: %s", typeName, body)
	}

	return nil
}