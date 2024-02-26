## GPT_FOFA
这是一个利用OpenAI转发API_key做请求（所以要去go-openai的config中修改相关数据），然后输入FOFA_API得到测绘结果的小Demo

### Result.xlsx
记录由FOFA搜索得到的结果

### fine_tune.go
如果需要用到大模型微调，则需要visa卡去购买openai的服务，这里给出代码示例

### pure_train_prepared.jsonl
纯净的训练集用于形成直接给FOFA的输入

### train_prepared.jsonl
普通的训练集用于FOFA相关的日常对话

### main.go
实现流交流和api访问的主逻辑
