# HCS but Gecko driver
심심해서 만든 셀레니움 & Gecko 드라이버를 통한 자가진단 자동화입니다.

진짜 할짓 없어서 만든거라 학교명과 지역 선택하는 부분이 하드코딩되어 있습니다.

다른 학교에서 사용하실 분이라면 개조해서 쓰시면 됩니다.

> 가끔씩 교육청이 이상한 인증서를 떤져주는데 그냥 무시하고 다시 하면 됩니다. (교육청 일 안하냐)

![](docs/Animation.gif)

사용법은 간단합니다
```
hcsbutgeckodriver 이름 생년월일 비밀번호
```

ex) `hcsbutgeckodriver 박민혁 050130 1234`

릴리즈 넣기 귀찮으니 그냥 zip 다운 받아서 쓰시면 됩니다.

윈도우64 빌드는 이미 되어있으나 다른 os나 arch에서 쓰실분은 `go build .` 하시면 됩니다
