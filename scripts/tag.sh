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

genGoCode="package version\n\nconst WDB_VERSION = \"$TARGET_VERSION\"\nconst CLI_VERSION = \"$CTL_VERSION\""

echo $genGoCode >$VERSION_GO_PATH

BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [[ "$BRANCH" != "main" ]]; then
    git restore .
    echo 'Aborting Commit, Tag, Push...'
    exit 1
fi

git add $VERSION_JSON_PATH $VERSION_GO_PATH
git commit -m "$COMMIT_MESSAGE"
git tag $TARGET_VERSION main
git push

echo "$TARGET_VERSION - Tag Released"
