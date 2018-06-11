### Introduction

This repository connect **wechat public account** and intercom.

When client ask question via wechat public account, the message will be forward to intercom. 

Then the company can reply client via intercom, the reply message will be forward to wechat client.

---

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

This two files need to change accrodingly.

src\intercom\intercom_setting.go

src\wechat\wechat_setting.go

---

### Run

```bash
export $GOPATH=`pwd`
go build src/main.go
sudo ./main.go 
```

---

### Related

the following repository connect **wechat personal account** and intercom:

https://github.com/richardchien/wechat-intercom
