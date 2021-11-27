// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package giteerepo

import (
	"context"
	"fmt"
	"gitee.com/openeuler/go-gitee/gitee"
	"sync"


	"github.com/ossf/scorecard/v3/clients"
	sce "github.com/ossf/scorecard/v3/errors"
)

type releasesHandler struct {
	client   *gitee.APIClient
	once     *sync.Once
	ctx      context.Context
	errSetup error
	owner    string
	repo     string
	releases []clients.Release
}

func (handler *releasesHandler) init(ctx context.Context, owner, repo string) {
	handler.ctx = ctx
	handler.owner = owner
	handler.repo = repo
	handler.errSetup = nil
	handler.once = new(sync.Once)
}

func (handler *releasesHandler) setup() error {
	handler.once.Do(func() {
		releases, _, err := handler.client.RepositoriesApi.GetV5ReposOwnerRepoReleases(
			handler.ctx, handler.owner, handler.repo, &gitee.GetV5ReposOwnerRepoReleasesOpts{})
		if err != nil {
			handler.errSetup = sce.WithMessage(sce.ErrScorecardInternal, fmt.Sprintf("githubv4.Query: %v", err))
		}
		handler.releases = releasesFrom(releases)
	})
	return handler.errSetup
}

func (handler *releasesHandler) getReleases() ([]clients.Release, error) {
	if err := handler.setup(); err != nil {
		return nil, fmt.Errorf("error during graphqlHandler.setup: %w", err)
	}
	return handler.releases, nil
}

func releasesFrom(data []gitee.Release) []clients.Release {
	var releases []clients.Release
	for _, r := range data {
		release := clients.Release{
			TagName:         r.TagName,
			//URL:             r.Body,
			TargetCommitish: r.TargetCommitish,

		}
		release.Assets = append(release.Assets,clients.ReleaseAsset{
			Name: r.Name,
			URL:  r.Assets,
		})
	}
	return releases
}
