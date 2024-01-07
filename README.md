# uuidconv
convert uuid to base64
## Cài đặt
```
go install github.com/hungtrd/uuidconv@latest
```

## Sử dụng
```
// new uuid
// mặc định sẽ copy dạng string vào clipboard
// -b64 để copy dạng base64 vào clipboard
uuidconv [-b64]

// encode
uuidconv -uuid 2183dc57-ddaf-4bb0-b861-5006ad3c4a29

// decode
uuidconv -uuid IYPcV92vS7C4YVAGrTxKKQ== -decode
```
