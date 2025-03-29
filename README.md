# it-support-service

ビジネスマッチングサービスを模したアプリ

## 技術選定

### バックエンド

Go を選定

- コンテナのイメージサイズを小さくしやすい
- 優れた後方互換性
- モノリスでもハマるし、将来マイクロサービス化したいとなった場合もスイッチングコストが最小限にしやすい
- (MVC のように決まった型なしのため)アーキテクチャが柔軟にしやすい

ORM は Eager Load, Upsert まで含めて一通りやりたいことが網羅されている SQLBoiler を、

Migration は up-down ができて、status が管理される sql-migrate を使用

### フロントエンド

Next.js(App Router)を採用

※ 候補として、React Router v7(フレームワーク)もあった

| フレームワーク         | pros                                                                                                                                                    | cons                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| Next.js(App Router)    | ・エコシステムが優秀で柔軟性あり<br>・クライアント依存のパフォーマンス低下を防げる<br>・ファイルベースルーティング<br>・モノレポ(Turborepo)と相性が良い | ・学習コスト<br>・状態管理が Remix と比べると煩わしくなる                          |
| React Router v7(Remix) | ・状態管理がスマートにしやすい<br>・パフォーマンスのクライアント依存を防ぎやすい                                                                        | ・Cookie 周りの実装が煩わしくなる<br>・モノレポの事例が Next.js と比べると限られる |

## スキーマ駆動

組織作りを始めるフェーズを想定し、REST API で

### インフラ

安さとスケーラビリティ重視

- Cloud Run
- Cloud Storage
- TiDB Serverless

## コマンド類

現状、各サービスごとのデプロイだが、後で整理したい

### マイグレーション

```
cd migrations && gcloud builds submit --config ./cloudbuild.yaml
```

### Registration API

```
cd api_server/registration && gcloud builds submit --config ./cloudbuild.yaml
```

### Business API

```
cd api_server/business && gcloud builds submit --config ./cloudbuild.yaml
```

### フロントエンド

```
cd frontend/it-support && gcloud builds submit --config ./cloudbuild.yaml
```

#### OpenAPI から型の自動生成

```
pnpx tsx generate_api_types.ts <openapi/yaml path> apps/{appName}/apis/generated/
```

### 参考

- Tidb へ Go からの接続
  - https://docs.pingcap.com/ja/tidb/stable/dev-guide-sample-application-golang-sql-driver/
  - https://docs.pingcap.com/ja/tidb/v6.5/dev-guide-choose-driver-or-orm/
  - https://zenn.dev/furegura/scraps/916150553fb2c5
- Go のプロジェクトレイアウト
  - https://github.com/golang-standards/project-layout/blob/master/README_ja.md
- sql-migrate を Cloud Run Job で実行
  - https://qiita.com/bayobayo0324/items/352d8bbb1bd7bcce8844
- Turborepo プロジェクトのビルド・デプロイ
  - https://turbo.hector.im/repo/docs/handbook/deploying-with-docker
  - https://turbo.build/repo/docs/guides/tools/docker
  - https://zenn.dev/simo_hr/articles/cbcd036c8814c3
  - https://zenn.dev/anneau/scraps/f2a2b6b9b0f387
  - https://nextjs.org/docs/pages/api-reference/config/next-config-js/output
- Next.js の Standalone ビルド
  - https://zenn.dev/rehabforjapan/articles/save-data-space-dockerfile
- Turborepo の packages/ui の Tailwind の使用
  - https://stackoverflow.com/questions/79416157/how-to-enable-tailwind-css-v4-0-for-the-packages-ui-components-in-turborepo
