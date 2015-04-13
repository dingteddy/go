###########################################
# build sub module
###########################################

GOSRCS := $(wildcard *.go)
GOES := $(patsubst %.go,%,$(GOSRCS))

all:
	@echo "start build..."
	go build $(GOSRCS)


.PHONY:clean rebuild
clean:
	@echo "start clean..."
	rm -f $(GOES)

rebuild:clean all
