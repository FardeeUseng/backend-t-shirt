# Project Name

```
dev
  ## สร้างไฟล์ air
    1. รันคอมมาน air init
    2. config file .air
      2.1 bin = "tmp\\main.exe"                  -> "./app/tmp/main.exe"
      2.2 cmd = "go build -o ./tmp/main.exe ."   -> "go build -o ./app/tmp/main.exe ./app/main.go"

  ## Dev Command
    air -c .air.toml -d  = Back Office API
    air -c .air.pdf.toml -d  = PDF Server
    air -c .air.test.toml -d = Back Office Test
```
