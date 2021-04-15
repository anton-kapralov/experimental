echo 'addends:1,addends:2' |
  protoc --encode kae.experimental.calculator.SumRequest ./proto/calculator/calculator_service.proto |
  curl -s --request POST \
    --header "Content-Type: application/protobuf" \
    --data-binary @- \
    http://localhost:8080/twirp/kae.experimental.calculator.CalculatorService/Sum |
  protoc --decode kae.experimental.calculator.SumResponse ./proto/calculator/calculator_service.proto
