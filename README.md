# golang_study

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

## 폴더 구조

```
server/
├── cmd/                    # 실행 파일 (main.go 등)
│   └── app/
│       └── main.go         # 애플리케이션 진입점
├── ent/                    # Ent 라이브러리에서 사용하는 폴더(스키마를 제외하곤 모두 자동 생성)
│   ├── schema/             # 스키마 정의
│   │   └── item.go         # 서비스 스키마 정의
├── internal/               # 내부 로직 (외부에 노출하지 않는 코드)
│   ├── api/                # API 핸들러 (HTTP 핸들러)
│   │   ├── router.go       # 라우터 설정 (e.g., Gorilla Mux, Echo 등)
│   │   └── item.go         # 서비스 관련 API 핸들러
│   ├── db/                 # 데이터베이스 로직
│   │   └── item.go         # 서비스 관련 DB 함수
│   ├── models/             # 데이터베이스 모델
│   │   └── item.go         # 서비스 관련 struct 정의
│   ├── service/            # 비즈니스 로직 (핸들러와 DB 로직 사이)
│   ├── config/             # 설정 파일 및 환경 변수 관리
│   │   ├── config.go       # 플래그 설정 관련 함수
│   │   └── database.go     # DB 연결 설정
│   └── utils/              # 공통 유틸리티 함수
│   │   ├── network.go      # 네트워크 관련 함수
│   │   └── path.go         # 경로 관련 함수
├── go.mod                  # Go 모듈 파일
└── README.md               # 프로젝트 설명
```

<br>

## Conventions

### 파일 네이밍

- 모두 소문자로 작성하고 밑줄로 여러 단어를 구분한다.
- "." 또는 "\_"로 시작하는 파일 이름은 제외된다.
- "\_test.go"로 끝나는 파일은 `go test`에서만 컴파일된다.

<br>

### 함수와 메서드

- 기본적으로 camel 케이스를 사용하고, 패키지 외부에 노출시켜야 한다면 대문자로 시작해야 한다.

<br>

### 상수

- 모두 대문자를 사용하고, 밑줄로 여러 단어를 구분한다.

<br>

### 변수

- 가능하다면 짧게 작성한다.
- 변수형이 `bool`인 경우, 변수명은 "Has", "Is", "Can", "Allow" 등으로 시작되어야 한다.
- i, j, k와 같은 단일 문자는 인덱스를 뜻한다.

<br>

## 목록

- [gen_pdf_file.go](./pdf/gen_pdf_file.go) : PDF 파일 생성하기
