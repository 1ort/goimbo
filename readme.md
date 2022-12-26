
<p align="center">
  <img width="192" height="192" src="https://user-images.githubusercontent.com/83316072/209536409-c691c252-3d12-4af2-93ab-e2278b9ff9b6.png">
</p>

# **Goimbo** is a Go-powered textboard / imageboard engine

![Screenshot 2022-12-26 at 13-40-32 _c_ - Board C](https://user-images.githubusercontent.com/83316072/209534091-c05c4d4f-02cd-49f5-9f91-ff2601a6168d.png)

# Setup
It is assumed that you already have Go and GIt installed, as well as the Postgres database.

1. Clone the repo
```bash
git clone https://github.com/1ort/goimbo.git
cd goimbo
```
2. Install dependencies
```bash
go mod download
```
3. Create and edit config file
```bash
cp base_config.yaml config.yaml
```
4. Build and run!
```bash
go build .
./goimbo
```

# To-Do

- Read-only API
- Post API using tokens
- CSRF protection
- Captcha
- Admin functionality
- Upload images and files
- Tripcodes
- password-based file/post deletion
- Anti-spam/anti-flood