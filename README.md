# 게시판 샘플

이 게시판은 3-tier 아키텍처를 기반으로 하며, 웹 프론트엔드, 백엔드 서버, 데이터베이스로 구성됩니다.

WebServer - WAS - DB 3-tier 아키텍처 간단하게 맛보자

## Frontend

### 화면별 기능 정의

#### Post 목록 화면 (/index.html)

- Post 목록을 표시합니다.
- 각 row를 클릭하면 해당 Post의 상세 정보 페이지로 이동합니다.

#### Post 상세 정보 화면 (/detail.html?id={id})

- 선택한 Post의 상세 정보를 표시합니다.
- 수정 또는 삭제 기능을 제공할 수 있습니다.

### 화면별 호출하는 API

#### Post 목록 화면 (/)

- `GET /post`: 모든 Post의 목록을 가져옵니다.

#### Post 상세 정보 화면 (/post/:id)

- `GET /post?id={id}`: 특정 ID의 Post의 상세 정보를 가져옵니다.

## Backend

### API 별 기능 정의

#### GET /post

- 모든 Post의 목록을 반환합니다.

#### GET /post?id={id}

- 특정 ID의 Post의 상세 정보를 반환합니다.

#### POST /post

- 새로운 Post를 추가합니다.
- 존재하는 Post의 경우 수정합니다.

#### DELETE /post?id={id}

- 특정 ID의 Post를 삭제합니다.

## Database

### Post Table 스키마 작성

#### Post Table

| Column      | Type       | Description             |
|-------------|------------|-------------------------|
| id          | INT        | Unique identifier       |
| title       | VARCHAR    | Post title              |
| content     | TEXT       | Post content            |
| createDate  | DATETIME   | Creation date of post   |
| updateDate  | DATETIME   | Last update date        |

### DDL 정리

```sql
CREATE TABLE Post (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT,
  createDate DATETIME DEFAULT CURRENT_TIMESTAMP,
  updateDate DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Network

### Webserver

- Nginx 웹서버 구성의 경우
```config
server {
    listen 80;
    server_name example.com; # 도메인 또는 IP 주소

    root /path/to/your/frontend; # 프론트엔드 파일이 위치한 디렉토리

    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:8080; # 백엔드 서버 주소
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```
