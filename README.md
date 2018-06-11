### Setup

```bash
git clone --recursive git@github.com:mayue42/wechat_intercom.git
```

or

```bash
git clone git@github.com:mayue42/wechat_intercom.git
cd wechat_intercom
git submodule init
git submodule update
```

or

```bash
go get gopkg.in/intercom/intercom-go.v2
pushd
mkdir -p src/github.com/pborman
cd src/github.com/pborman
git clone https://github.com/pborman/uuid.git
popd
```

---

### Settings

src\intercom\intercom_setting.go

src\wechat\wechat_setting.go

---

### Run

```bash
export $GOPATH=`pwd`
go build src/main.go
sudo ./main.go 
```
