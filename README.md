# vidflow
- vidflow是一个可以自动调用gpt接口，如openAi, qwen等接口，拿到一段古诗词
- 拿到古诗词后继续调用文生图模型，生成图片
- 拿到图片后调用语音模型与视频模型，生成一段视频
- 将视频上传到抖音，快手等短视频平台。

# DDD

## 上下文

Poem Generation（诗词生成上下文）
职责：生成一段古诗词，并保存相关的诗词信息。
边界：调用外部 GPT 接口（如 OpenAI 或其他语言模型）。  

Poem Generation 聚合

聚合根：Poem（诗词实体）
属性：
ID（唯一标识符）
Content（诗词内容）
Topic（主题）
CreatedAt（创建时间）


Image Generation（图像生成上下文）\
职责：根据诗词内容生成图像。生成的图像可能需要与诗词一一关联。\
边界：调用外部文生图模型接口，并保存图像信息。  
聚合根
- image



Video Generation（视频生成上下文）\
职责：根据诗词和生成的图像合成视频，加入语音和背景音乐。\
边界：调用外部语音模型和视频合成服务。

Video Distribution（视频分发上下文）\
职责：将生成的视频上传到抖音、快手等短视频平台。\
边界：与分发平台（抖音、快手）交互，记录上传状态。


+----------------------+        +-----------------------+        +-----------------------+        +-----------------------+
| Poem Context         | ----> | Image Context         | ----> | Video Context         | ----> | Distribution Context  |
+----------------------+        +-----------------------+        +-----------------------+        +-----------------------+
| Poem                |        | Image                |        | Video                |        | DistributedVideo      |
| - ID                |        | - ID                |        | - ID                |        | - ID                 |
| - Content           |        | - URL               |        | - URL               |        | - Platform           |
| - Topic             |        | - PoemID            |        | - PoemID            |        | - Upload Status      |
| - CreatedAt         |        | - CreatedAt         |        | - Status            |        | - Upload Time        |
+----------------------+        +-----------------------+        +-----------------------+        +-----------------------+
