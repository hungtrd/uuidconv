# uuidconv
Chuyển đổi giữa các format UUID

web version: [hungtrd.github.io/uuidconv](https://hungtrd.github.io/uuidconv)
## Cài đặt
```
go install github.com/hungtrd/uuidconv@latest
```

## Sử dụng
```
// Xem hướng dẫn sử dụng
uuidconv -h
```
```
// Tạo mới uuid
uuidconv

// Tạo mới và copy vào clipboard
uuidconv -c base62

// Chuyển đổi giữa các format
uuidconv convert -f string -t base64 2183dc57-ddaf-4bb0-b861-5006ad3c4a29

// Hoặc có thể không cần chỉ định loại format input
uuidconv convert -t base64 2183dc57-ddaf-4bb0-b861-5006ad3c4a29
```
