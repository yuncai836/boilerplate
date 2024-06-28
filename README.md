## 工具包

配置，可观测组件，错误，日志库的样板代码

### 配置初始化

- 可以从环境变量获取 apollo 配置，然后通过配置中心获取其他配置。
- 可以从环境变量获取 本地配置文件路径，然后通过 viper 读取。

#### apollo 环境变量

- apollo_addr
- apollo_app_id
- apollo_cluster
- apollo_namespace
- apollo_secret

#### 本地配置环境变量

- config_path