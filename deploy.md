#### 准备密钥文件

将主网节点生成的密钥文件放到 config 目录下。例如：sequencer.keystore

#### 准备配置文件

将 config 目录下的 config.local.toml 文件复制一份并改名。例如：bridge.config.toml

修改 bridge.config.toml 中的配置

* 修改 ClaimTxManager 下的 PrivateKey 的 Password 为 主网节点生成密钥文件时的密码，以便在程序中能打开并读取 sequencer.keystore 文件。例如：Ryoshi
* 修改 Etherman 下的 L2URLs 为 L2 的地址。例如：http://3.91.48.119:8123
* 修改 Etherman 下的 L2ChainIds 为 L2 的编号。例如：11198111
* 修改 Etherman 下的 L1URL 为 L1 的地址。例如：http://sepolia.shibchain.io:8545
* 修改 Etherman 下的 L1ChainId 为 L1 的编号。例如：11155111
* 修改 NetworkConfig 下的 GenBlockNumber 为 L1 部署合约时的区块。例如：5575687
* 修改 NetworkConfig 下的 PolygonBridgeAddress 为 L1 上部署的桥合约的地址。例如：0x7874aae070d58702dFd7B81e3c3Eed774665c836
* 修改 NetworkConfig 下的 PolygonZkEVMGlobalExitRootAddress 为 L1 上部署的全局退出根合约的地址。例如：0x3F7C86731D3cf980Fd7a090A6535E4E5e5eF1101
* 修改 NetworkConfig 下的 L2PolygonBridgeAddresses 为 L2 上部署的桥合约的地址。例如：0x7874aae070d58702dFd7B81e3c3Eed774665c836


#### 修改 docker-compose.yml 文件
修改 x1-bridge-service 服务的 volumes， 将上述两个文件映射到镜像内。例如：

```text
./config/sequencer.keystore:/pk/keystore.claimtxmanager
./config/bridge.config.toml:/app/config.toml
```
#### 构建镜像

```shell
make build-docker
```

#### 修改 Makefile 文件

修改 run 项，移除不需要启动的服务。
改为:

```text
run: stop ## runs all services
	$(RUN_DBS)
	$(RUN_REDIS)
	$(RUN_ZOOKEEPER)
	sleep 3
	$(RUN_COIN_KAFKA)
	sleep 3
	$(RUN_BRIDGE)
```

#### 启动服务

```shell
make run
```
