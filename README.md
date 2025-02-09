# it-support-service
ビジネスマッチングサービスを模したアプリ

## 技術選定
### バックエンド
Goを選定
- コンテナのイメージサイズを小さくしやすい
- 優れた後方互換性
- モノリスでもハマるし、将来マイクロサービス化したいとなった場合もスイッチングコストが最小限にしやすい
- (MVCのように決まった型なしのため)アーキテクチャが柔軟にしやすい

ORMはEager Load, Upsertまで含めて一通りやりたいことが網羅されているSQLBoilerを、

Migrationはup-downができて、statusが管理されるsql-migrateを使用

### フロントエンド
Next.js(App Router)を採用

※ 候補として、React Router v7(フレームワーク)もあった

| フレームワーク | pros | cons |
| ------------- | ------------- | ------------- |
| Next.js(App Router) | ・エコシステムが優秀で柔軟性あり<br>・クライアント依存のパフォーマンス低下を防げる<br>・ファイルベースルーティング<br>・モノレポ(Turborepo)と相性が良い | ・学習コスト<br>・状態管理がRemixと比べると煩わしくなる |
| React Router v7(Remix) | ・状態管理がスマートにしやすい<br>・パフォーマンスのクライアント依存を防ぎやすい | ・Cookie周りの実装が煩わしくなる<br>・モノレポの事例がNext.jsと比べると限られる |

## スキーマ駆動
組織作りを始めるフェーズを想定し、REST APIで

### インフラ
安さとスケーラビリティ重視
- Cloud Run
- Cloud Storage
- TiDB Serverless
