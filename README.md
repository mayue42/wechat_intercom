### set up:

go get gopkg.in/intercom/intercom-go.v2

pushd src/gopkg.in/intercom/intercom-go.v2

git remote add bugfix git@github.com:harryhare/intercom-go.git

git pull bugfix v2

popd

mkdir -p src/github.com/pborman


### settings:

src\intercom\intercom_setting.go

src\wechat\wechat_setting.go

