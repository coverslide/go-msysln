project=msysln
postfix=.exe

executable=$(project)$(postfix)

$(executable):
	go build -o $(executable) main.go

clean:
	-rm $(executable)

.PHONY: clean
