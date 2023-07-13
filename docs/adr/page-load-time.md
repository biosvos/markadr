# 웹페이지 로딩 속도 향상

- Status: accepted

## Context and Problem Statement

/pages/{파일} 웹페이지 로딩이 너무 느리다.
빠르게 웹사이트를 로딩하고 싶다.

with disable cache
index 현재: 238ms, 496ms, 78ms
page 현재: 549ms, 967ms, 1.04s, 792ms

## Decision Drivers

- page 웹 로딩 시간

## Options

### 페이지 결과 캐싱

markdown 파일을 읽어서 파싱하는 과정 때문에 느린가? 싶어서

| Pros                | Cons                       |
|---------------------|----------------------------|
| 결과를 캐싱하면 아무래도 빨라진다. | 파일 내용이 변경될 때에 대한 처리가 필요하다. |

### 폰트 프리로드

| Pros                | Cons |
|---------------------|------|
| 캐싱이랑 비교하면 신경 쓸게 없다. |      |

### repository에 폰트 저장

| Pros                | Cons                 |
|---------------------|----------------------|
| 캐싱이랑 비교하면 신경 쓸게 없다. | repository 사이즈가 커진다. |

## Decision Outcome

Pick: 폰트 프리로드

- 물론, disabled cache를 껐지만, 100ms 이하로 걸린다.

## Links

- [Link type](link to adr)
