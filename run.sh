if [[ $1 = "build" ]]; then
  go build -o homework main.go user_controller.go user_model.go utils.go form.go
fi

if [[ $1 = "test" ]]; then
    go test
fi

if [[ $1 = "" ]]; then
    go run main.go user_controller.go user_model.go utils.go form.go
fi