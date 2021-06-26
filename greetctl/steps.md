1. go mod init greetctl (Create a go module for CLI.)
2. go get -u github.com/spf13/cobra/cobra (Get Cobra library)
3. cobra init --pkg-name greetctl (Create a bare minimum skeleton)
4. go get github.com/mitchellh/go-homedir@v1.1.0 (library required to get home-dir without cgo)
5. run the command go run main.go to test.
6. To add a new command do cobra add create. New go file will be created under ./cmd/create.go
7. Do cobra add card, By default, both create and card commands will be added to the top-level command (root.go). But we want to add card (card.go) sub-command to create (create.go) command instead of top-level command (root.go) so lets edit /cmd/card.go. Add cardCmd in createCmd instead of rootCmd. </br>
![initial](./img/initial-card.pngimage.png)
