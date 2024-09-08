![](https://img.shields.io/badge/rpm_x86_build-success-brightgreen) &nbsp;
![](https://img.shields.io/badge/rpm_arm_build-success-brightgreen) &nbsp;
![](https://img.shields.io/badge/deb_arm_build-success-brightgreen) &nbsp;
![](https://img.shields.io/badge/deb_amd_build-success-brightgreen)&nbsp;
<br>

<div align="center">
  <img src="https://github.com/user-attachments/assets/0429e5e9-59e1-407f-9ddd-9c3b8d2f9fdb" alt="mark1 logo" width="150"/>
</div>

# mark1🌱

<details>
<summary>日本語</summary>

### 概要
このシステムはシンプルな認証基盤を簡単に構築し、提供するためのものです。

### インストール
[Releases](https://github.com/ryo-arima/mark1/releases) からファイルをダウンロードして、インストールしてください。

##### For Red Hat-based Systems (RHEL, CentOS, Fedora) -- aarch64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-aarch64.aarch64.rpm
$ sudo rpm -ivh mark1-0.0.1-aarch64.aarch64.rpm
```
##### For Red Hat-based Systems (RHEL, CentOS, Fedora) -- x86_64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-x86_64.x86_64.rpm
$ sudo rpm -ivh mark1-0.0.1-x86_64.x86_64.rpm
```

##### For Debian-based Systems (Debian, Ubuntu) -- arm64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1_0.0.1_arm64.deb
$ sudo dpkg -i mark1_0.0.1_arm64.deb
```

##### For Debian-based Systems (Debian, Ubuntu) -- amd64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1_0.0.1_amd64.deb
$ sudo dpkg -i mark1_0.0.1_amd64.deb
```

### 設定
※MysqlとRedisが必要なので、事前にインストール、設定、起動してください<br>
このレポジトリの`docker-compose.yml`を使って起動することもできます。<br>
<br>
設定ファイルは `/etc/mark1/main.yaml` にあります。このファイルを編集してください。
<details><summary>main.yaml</summary>

```yaml
Application:
  Common:
    MemberRoles:
      - group_admin
      - group_editor
      - group_viewer
  Server:
    port: 8000
    admin:
      emails:
        - "admin@mail.com"
    jwt:
      key: "secret"
    mail:
      host: "smtp.mail.com"
      port: 587
      user: ""
    tmp:
      letters: ""
      length: 6
  Client:
    HomePath: "etc/.mark1"
    ServerEndpoint: "http://localhost:8000"
    UserEmail: "user1@mail.com"
    UserPassword: "secret"
    HomeDir: "etc/.mark1"

MySQL:
  host: 127.0.0.1
  user: root
  pass: mysql
  port: 3306
  db: mark1

Redis:
  host: 127.0.0.1
  port: 6379
  user: "default"
  pass: "mysecretpassword"
  db: 0
```
##### クライアントの役割を`$HOME/.mark1/role`に設定する場合
```bash
cat $HOME/.mark1/role ### adminを設定する場合は、main.yamlのApplication.Server.admin.emailsにadminのメールアドレスを追加してください
admin
```

</details>

### 使い方

##### サーバーを起動します
```bash
$ sudo systemctl start mark1
```

##### クライアントを実行します
```bash
$ mark1 --help
Usage:
  mark1-admin [command]

Available Commands:
  bootstrap   bootstrap the value of a key
  create      create the value of a key
  delete      delete the value of a key
  get         get the value of a key
  help        Help about any command
  update      update the value of a key

Flags:
  -h, --help   help for mark1-admin

Use "mark1-admin [command] --help" for more information about a command.
```

## 🤝 貢献
自由です。荒らさない程度に自由に貢献できます。バックアップもあるので、心配しないでください。
皆さんの力を貸してください。🙏

## 📞 お問い合わせ

質問や問題があれば、イシューに書き込むか[メール](mailto:ryo.arima.zzz@gmail.com)してください。

</details>
<details>
<summary>English</summary>

### Overview
This system is designed to easily build and provide a simple authentication foundation.

### Installation
Download the file from [Releases](https://github.com/ryo-arima/mark1/releases)

##### For Red Hat-based Systems (RHEL, CentOS, Fedora) -- aarch64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-aarch64.aarch64.rpm
$ sudo rpm -ivh mark1-0.0.1-aarch64.aarch64.rpm
```
##### For Red Hat-based Systems (RHEL, CentOS, Fedora) -- x86_64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1-0.0.1-x86_64.x86_64.rpm
$ sudo rpm -ivh mark1-0.0.1-x86_64.x86_64.rpm
```

##### For Debian-based Systems (Debian, Ubuntu) -- arm64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1_0.0.1_arm64.deb
$ sudo dpkg -i mark1_0.0.1_arm64.deb
```

##### For Debian-based Systems (Debian, Ubuntu) -- amd64
```bash
$ wget https://github.com/ryo-arima/mark1/releases/download/v0.0.1/mark1_0.0.1_amd64.deb
$ sudo dpkg -i mark1_0.0.1_amd64.deb
```

### Configuration
* Mysql and Redis are required, so please install, configure, and start them in advance.<br>
You can also start using the `docker-compose.yml` in this repository.<br>
<br>
The configuration file is located at `/etc/mark1/main.yaml`. Please edit this file.
<details><summary>main.yaml</summary>

```yaml
Application:
  Common:
    MemberRoles:
      - group_admin
      - group_editor
      - group_viewer
  Server:
    port: 8000
    admin:
      emails:
        - "admin@mail.com"
    jwt:
      key: "secret"
    mail:
      host: "smtp.mail.com"
      port: 587
      user: ""
    tmp:
      letters: ""
      length: 6
  Client:
    HomePath: "etc/.mark1"
    ServerEndpoint: "http://localhost:8000"
    UserEmail: "user1@mail.com"
    UserPassword: "secret"
    HomeDir: "etc/.mark1"

MySQL:
  host: 127.0.0.1
  user: root
  pass: mysql
  port: 3306
  db: mark1

Redis:
  host: 127.0.0.1
  port: 6379
  user: "default"
  pass: "mysecretpassword"
  db: 0
```
</details>

##### If you want to set the role of the client to `$HOME/.mark1/role`
```bash
cat $HOME/.mark1/role ### If you want to set admin, please add the email address of admin to Application.Server.admin.emails in main.yaml
admin
```

### Usage

##### Start the server
```bash
$ sudo systemctl start mark1
```

##### Run the client
```bash
$ mark1 --help
Usage:
  mark1-admin [command]

Available Commands:
  bootstrap   bootstrap the value of a key
  create      create the value of a key
  delete      delete the value of a key
  get         get the value of a key
  help        Help about any command
  update      update the value of a key

Flags:
  -h, --help   help for mark1-admin

Use "mark1-admin [command] --help" for more information about a command.
```

## 🤝 Contributions
Free. You are free to contribute as long as you don't troll. Don't worry, we have backups.
We need your help.

## 📞 Contact
If you have any questions or problems, please post them in the issue or [email](mailto:ryo.arima.zzz@gmail.com).
</details>