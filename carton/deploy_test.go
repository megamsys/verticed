/*
** copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package carton

import (
	//"gopkg.in/check.v1"
)

/*func (s *S) TestDeployBox(c *check.C) {
	a := App{
		Name:     "someApp",
		Platform: "django",
		Teams:    []string{s.team.Name},
	}
	err := s.conn.Apps().Insert(a)
	c.Assert(err, check.IsNil)
	defer s.conn.Apps().Remove(bson.M{"name": a.Name})
	s.provisioner.Provision(&a)
	defer s.provisioner.Destroy(&a)
	writer := &bytes.Buffer{}
	err = Deploy({
		App:          &a,
		Version:      "version",
		Commit:       "1ee1f1084927b3a5db59c9033bc5c4abefb7b93c",
		OutputStream: writer,
	})
	c.Assert(err, check.IsNil)
	logs := writer.String()
	c.Assert(logs, check.Equals, "Git deploy called")
}



func (s *S) TestDeployToProvisioner(c *check.C) {
	a := App{
		Name:     "someApp",
		Platform: "django",
		Teams:    []string{s.team.Name},
	}
	err := s.conn.Apps().Insert(a)
	c.Assert(err, check.IsNil)
	defer s.conn.Apps().Remove(bson.M{"name": a.Name})
	s.provisioner.Provision(&a)
	defer s.provisioner.Destroy(&a)
	writer := &bytes.Buffer{}
	opts := {App: &a, Version: "version"}
	_, err = deployToProvisioner(&opts, writer)
	c.Assert(err, check.IsNil)
	logs := writer.String()
	c.Assert(logs, check.Equals, "Git deploy called")
}

func (s *S) TestDeployToProvisionerImage(c *check.C) {
	a := App{
		Name:     "someApp",
		Platform: "django",
		Teams:    []string{s.team.Name},
	}
	err := s.conn.Apps().Insert(a)
	c.Assert(err, check.IsNil)
	defer s.conn.Apps().Remove(bson.M{"name": a.Name})
	s.provisioner.Provision(&a)
	defer s.provisioner.Destroy(&a)
	writer := &bytes.Buffer{}
	opts := {App: &a, Image: "my-image-x"}
	_, err = deployToProvisioner(&opts, writer)
	c.Assert(err, check.IsNil)
	logs := writer.String()
	c.Assert(logs, check.Equals, "Image deploy called")
}

*/