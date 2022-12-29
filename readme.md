
<p align="center">
  <img width="192" height="192" src="https://user-images.githubusercontent.com/83316072/209536409-c691c252-3d12-4af2-93ab-e2278b9ff9b6.png">
</p>

# **Goimbo** is a Go-powered textboard / imageboard engine
| Board View  | Thread view |
| ------------- | ------------- |
| ![Board](https://user-images.githubusercontent.com/83316072/209942828-d681d675-f0b6-4e42-b90a-86beef8a21a0.png)  | ![Thread](https://user-images.githubusercontent.com/83316072/209942437-0d63c4ce-54d4-41aa-9b0e-dbacbe45902e.png) |



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
- [x] Read-only API
- [x] CSRF protection
- [x] Captcha
- [ ] Post API using tokens
- [ ] Admin functionality
- [ ] Upload images and files
- [ ] Tripcodes
- [ ] password-based file/post deletion
- [ ] Anti-spam/anti-flood
- [ ] Docker & docker-compose
- [ ] Dashboard
- [ ] Feedback form
