# Shell script to generate wdb version.go and release tag
#
# Usage: sh ./scripts/tag.sh 1.1.4
#        Do not add 'v' before the tag, it is added in the script itself.
#        Generates version.go file, git commit, tags and pushes only if current branch is main

rx='^([0-9]+\.){0,2}(\*|[0-9]+)$'
if [[ $1 =~ $rx ]]; then
    echo "Creating Tag..."
else
    echo "ERROR:not semver compliant: '$1'"
    exit 1
fi

if [[ $1 != v* ]]; then
    TARGET_VERSION=v$1
else
    TARGET_VERSION=$1
fi

VERSION_JSON_PATH="../internal/version/version.json"
VERSION_GO_PATH="../internal/version/version.go"
COMMIT_MESSAGE="pre-tag: updated wdb version: $TARGET_VERSION"

parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

sh $parent_path/upgrade.sh $TARGET_VERSION

BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [[ "$BRANCH" != "main" ]]; then
    echo 'Aborting Commit, Tag, Push...'
    git restore $VERSION_JSON_PATH $VERSION_GO_PATH
    exit 1
fi

go fmt ./...

git add $VERSION_JSON_PATH $VERSION_GO_PATH
git commit -m "$COMMIT_MESSAGE"
git push

git tag $TARGET_VERSION main
git push --tags

echo "$TARGET_VERSION - Tag Released"
