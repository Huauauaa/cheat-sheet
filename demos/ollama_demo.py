from openai import OpenAI

# 关键：将 base_url 指向本地的 Ollama 服务
client = OpenAI(
    base_url='http://localhost:11434/v1',
    api_key='ollama' # 可填写任意非空字符串，Ollama不验证
)

# 像调用 OpenAI API 一样调用
response = client.chat.completions.create(
    model='llama3.2', # 使用你本地的模型名
    messages=[
        {'role': 'user', 'content': '你好，请自我介绍。'}
    ]
)

print(response.choices[0].message.content)