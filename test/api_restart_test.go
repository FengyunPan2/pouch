package main

import (
	"net/url"

	"github.com/alibaba/pouch/test/environment"
	"github.com/alibaba/pouch/test/request"

	"github.com/go-check/check"
)

// APIContainerRestartSuite is the test suite for container upgrade API.
type APIContainerRestartSuite struct{}

func init() {
	check.Suite(&APIContainerRestartSuite{})
}

// SetUpTest does common setup in the beginning of each test.
func (suite *APIContainerRestartSuite) SetUpTest(c *check.C) {
	SkipIfFalse(c, environment.IsLinux)
}

// TestAPIContainerRestart is to verify restarting container.
func (suite *APIContainerRestartSuite) TestAPIContainerRestart(c *check.C) {
	cname := "TestAPIContainerRestart"

	CreateBusyboxContainerOk(c, cname)

	resp, err := request.Post("/containers/" + cname + "/start")
	c.Assert(err, check.IsNil)
	CheckRespStatus(c, resp, 204)

	q := url.Values{}
	q.Add("t", "1")
	query := request.WithQuery(q)

	resp, err = request.Post("/containers/"+cname+"/restart", query)
	c.Assert(err, check.IsNil)
	CheckRespStatus(c, resp, 204)

	DelContainerForceOk(c, cname)
}

// TestAPIRestartStoppedContainer it to verify restarting a stopped container.
func (suite *APIContainerRestartSuite) TestAPIRestartStoppedContainer(c *check.C) {
	cname := "TestAPIContainerRestart"

	CreateBusyboxContainerOk(c, cname)

	q := url.Values{}
	q.Add("t", "1")
	query := request.WithQuery(q)

	resp, err := request.Post("/containers/"+cname+"/restart", query)
	c.Assert(err, check.IsNil)
	CheckRespStatus(c, resp, 500)

	DelContainerForceOk(c, cname)
}
