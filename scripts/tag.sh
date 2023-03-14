# Shell script to generate wdb version.go and release tag
#
# Usage: sh ./scripts/tag.sh 1.1.4
#        Do not add 'v' before the tag, it is added in the script itself.
#        Generates version.go file, git commit, tags and pushes only if current branch is main

TARGET_VERSION=$1

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

git add $VERSION_JSON_PATH $VERSION_GO_PATH
git commit -m "$COMMIT_MESSAGE"
git push

git tag $TARGET_VERSION main
git push --tags

echo "$TARGET_VERSION - Tag Released"
