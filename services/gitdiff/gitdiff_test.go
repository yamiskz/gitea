	jsoniter "github.com/json-iterator/go"
			got, err := ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(testcase.gitdiff))
			json := jsoniter.ConfigCompatibleWithStandardLibrary
	result, err := ParsePatch(20, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(40, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(40, 5, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(20, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(40, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff2))
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff2a))
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff3))
	assert.NoError(t, models.PrepareTestDatabase())
	issue := models.AssertExistsAndLoadBean(t, &models.Issue{ID: 2}).(*models.Issue)
	user := models.AssertExistsAndLoadBean(t, &models.User{ID: 1}).(*models.User)
		diffs, err := GetDiffRangeWithWhitespaceBehavior(gitRepo, "559c156f8e0178b71cb44355428f24001b08fc68", "bd7063cc7c04689c4d082183d32a604ed27a24f9",
			setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffFiles, behavior)
	diffRecord := diffMatchPatch.DiffMain(highlight.Code("main.v", "		run()\n"), highlight.Code("main.v", "		run(db)\n"), true)
		ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(testcase.gitdiff))