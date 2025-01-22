NAME=blog.local

dev: go.mod
	(./bin/air&) && cd render && npm run dev

all: go.mod
	go build && (./blog.local&) 

re: fclean all

go.mod:
	go mod init $(NAME)
	go mod tidy
	cd render && npm i && npm run build


fclean: clean
	go clean -cache
	go clean -modcache
	rm -rf go.mod go.sum
	rm -rf tmp

clean:
	rm -rf $(NAME)


.PHONY: all dev re fclean clean
