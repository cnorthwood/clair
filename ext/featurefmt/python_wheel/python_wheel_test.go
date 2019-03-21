// Copyright 2019 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package python_wheel

import (
	"testing"

	"github.com/coreos/clair/database"
	"github.com/coreos/clair/ext/featurefmt"
	"github.com/coreos/clair/ext/versionfmt/pep440"
)

func TestPythonEggFeatureDetection(t *testing.T) {
	for _, test := range []featurefmt.TestCase{
		{
			"valid case",
			map[string]string{
				"usr/lib/python/site-packages/foo.dist-info/METADATA": "python_wheel/testdata/metadata.txt",
				"usr/lib/python/site-packages/bar.dist-info/METADATA": "python_wheel/testdata/invalid.txt",
				"usr/lib/python/site-packages/baz.dist-info/METADATA": "python_wheel/testdata/missing-name.txt",
			},
			[]database.LayerFeature{
				{Feature: database.Feature{"buildbot-www", "1.1.0", "pep440", "source"}},
			},
		},
	} {
		featurefmt.RunTest(t, test, lister{}, pep440.ParserName)
	}
}
