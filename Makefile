###########################################
# top level build method
###########################################
MKDIR = mkdir
RM = rm

BIN_DIR = bin

# 需要排除的目录
exclude_dirs := include $(BIN_DIR)

# 取得当前子目录深度为1的所有目录名称
dirs := $(shell find . -maxdepth 1 -type d)
dirs := $(basename $(patsubst ./%,%,$(dirs)))
dirs := $(filter-out $(exclude_dirs),$(dirs))

# 避免clean子目录操作同名，加上_clean_前缀
SUBDIRS := $(dirs)

#
.PHONY: subdirs $(SUBDIRS) clean distclean rebuild

all:subdirs
	$(MKDIR) -p $(BIN_DIR)
	go build *.go

subdirs: $(SUBDIRS)

# 执行默认make
$(SUBDIRS): 
	$(MAKE) -C $@

# 执行clean
clean:
	for i in $(SUBDIRS); do $(MAKE) -C $$i $@; done

# 执行distclean
distclean:
	for i in $(SUBDIRS); do $(MAKE) -C $$i $@; done
	$(RM) -rf $(BIN_DIR)

rebuild:distclean all
