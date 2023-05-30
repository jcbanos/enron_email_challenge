build_indexer:
	go build indexer_f/indexer.go

run_indexer:
	./indexer enron_mail_20110402 localhost

build_server:
	go build server.go

run_server:
	./server localhost

install_zinc:
	mkdir -p zinc
	wget https://github.com/zinclabs/zincobserve/releases/download/v0.4.5/zincobserve-v0.4.5-linux-amd64.tar.gz
	tar -xvzf zincobserve-v0.4.5-linux-amd64.tar.gz
	mv zincobserve zinc/zincobserve
	rm zincobserve-v0.4.5-linux-amd64.tar.gz

make run_zinc:
	cd zinc
	ZO_ROOT_USER_EMAIL=root@example.com ZO_ROOT_USER_PASSWORD=Complexpass#123 zinc/zincobserve