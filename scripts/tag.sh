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
COMMIT_MESSAGE="pre-tag-commit: updated wdb version: $TARGET_VERSION"

echo "TARGET VERSION: $TARGET_VERSION"

parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

if [[ -z "$TARGET_VERSION" ]]; then
    # $var is empty, do what you want
    echo "version not provided"
fi

cd "$parent_path"
WDB_VERSION=$(cat $VERSION_JSON_PATH | jq -r ".wdb_version")
CTL_VERSION=$(cat $VERSION_JSON_PATH | jq -r ".wdbctl_version")
WDB_BUILD_DATE=$(cat $VERSION_JSON_PATH | jq -r ".wdb_build_date")

TARGET_BUILD_DATE=$(date +%F_%T)

genGoCode="package version\n\nconst (\n\tWDB_VERSION = \"$TARGET_VERSION\"\n\tCLI_VERSION = \"$CTL_VERSION\"\n\tBUILD_DATE = \"$TARGET_BUILD_DATE\"\n)"

tmp=$(mktemp)
jq --arg v "$TARGET_VERSION" '.wdb_version = $v' $VERSION_JSON_PATH >"$tmp" && mv "$tmp" $VERSION_JSON_PATH
jq --arg b "$TARGET_BUILD_DATE" '.wdb_build_date = $b' $VERSION_JSON_PATH >"$tmp" && mv "$tmp" $VERSION_JSON_PATH
rm "$tmp"

echo $genGoCode > $VERSION_GO_PATH

BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [[ "$BRANCH" != "main" ]]; then
    echo 'Aborting Commit, Tag, Push...'
    git restore $VERSION_JSON_PATH $VERSION_GO_PATH
    exit 1
fi

# go fmt the version.go file
go fmt $VERSION_GO_PATH

# git add and commit tagged and built files
git add $VERSION_JSON_PATH $VERSION_GO_PATH
git commit -m "$COMMIT_MESSAGE"
git push

git tag $TARGET_VERSION main
git push --tags

echo "$TARGET_VERSION - Tag Released"
