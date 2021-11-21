## go mod commands
- Initiate project\
  `  go mod init  `
- Add dependecny \
  `go get <libsName>`
- Install in vendor (install all packages in local machine)\
`go mod vendor`
- Remove unused modules\
`go mod tidy`
  - Get un used modules list\
  `go mod tidy -v`
- Build\
` go build`
- Download the missing dependencies without having to run/build your project\
`go mod download`
- Verify go modules\
`go mod verify`
- Get list of modules available in cache\
`go list all`
  - Get list of modules used in project\
  `go list -m all`
  - Get list of perticular package version\
  `go list -m versions github.com/<package dir>`
- Edit mod file using commands\
 `go mod edit -go <version name>`\
 `go mod edit -module <module name>`
  
