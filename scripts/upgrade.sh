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
jq --arg v "$TARGET_VERSION" '.wdb_version = $v' $VERSION_JSON_PATH &&
    jq --arg b "$TARGET_BUILD_DATE" '.wdb_build_date = $b' $VERSION_JSON_PATH >"$tmp" >"$tmp" &&
    mv "$tmp" $VERSION_JSON_PATH

echo $genGoCode >$VERSION_GO_PATH

cd ../
go fmt ./...
