### Concurrency
Concurrency dealing với các tác vụ trong 1 khoảng thời gian
- điều phối nhiều tác vụ trong cùng 1 khoảng t/gian
- trong quá trình điều phối chỉ cho phép 1 tác vụ chạy trong 1 thời điểm

###
VD: trong 1 core, có 3 tác vụ A, B, C
xử lý A -> tạm dừng -> B -> switch A -> switch C -> switch B ...

### Parallelism
- thực thi nhiều task cùng 1 thời điểm (cpu core >= 2)
- doing lots of thing at once

### Thread vs Goroutine
- Thread được quản lý bởi kernal và phụ thuộc phần cứng
Goroutine được quản lý bởi go runtime

- Stack size của thread = 1M | gor = 8KB -> 1GB

- Giao tiếp giữa các gor qua channel

