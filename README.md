# go_study

Golang 공부 모음.zip

<br>

## Go 명령어

```bash
# 버전 확인
> go version

# go.mod 초기화 및 go.mod 파일 생성
> go mod init
```

<br>

## Conventions

### 파일 네이밍

- 모두 소문자로 작성하고 밑줄로 여러 단어를 구분한다.
- "." 또는 "\_"로 시작하는 파일 이름은 제외된다.
- "\_test.go"로 끝나는 파일은 `go test`에서만 컴파일된다.

### 함수와 메서드

- 기본적으로 camel 케이스를 사용하고, 패키지 외부에 노출시켜야 한다면 대문자로 시작해야 한다.

### 상수

- 모두 대문자를 사용하고, 밑줄로 여러 단어를 구분한다.

### 변수

- 가능하다면 짧게 작성한다.
- 변수형이 `bool`인 경우, 변수명은 "Has", "Is", "Can", "Allow" 등으로 시작되어야 한다.
- i, j, k와 같은 단일 문자는 인덱스를 뜻한다.

<br>

## 목록

- [gen_pdf_file.go](./pdf/gen_pdf_file.go) : PDF 파일 생성하기
