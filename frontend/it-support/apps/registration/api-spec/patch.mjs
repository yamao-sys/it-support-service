import path from "path";
import process from "process";
import * as fs from "fs";
import yaml from "js-yaml";

// OpenAPI YAML ファイルを読み込み
const openapiYamlPath = path.resolve(process.cwd(), "tsp-output", "schema", "openapi.yaml");
const input = fs.readFileSync(openapiYamlPath, "utf8");
const doc = yaml.load(input);

// encoding.contentType: "text/plain" をすべて削除
for (const pathKey in doc.paths) {
  const pathItem = doc.paths[pathKey];
  for (const methodKey in pathItem) {
    const operation = pathItem[methodKey];
    const content = operation?.requestBody?.content?.["multipart/form-data"];
    const encoding = content?.encoding;

    if (encoding && typeof encoding === "object") {
      for (const field in encoding) {
        if (encoding[field]?.contentType === "text/plain") {
          delete encoding[field].contentType;
        }

        // フィールドが空になったら削除（任意）
        if (Object.keys(encoding[field]).length === 0) {
          delete encoding[field];
        }
      }

      // encoding 全体が空になったら削除（任意）
      if (Object.keys(encoding).length === 0) {
        delete content.encoding;
      }
    }
  }
}

const output = yaml.dump(doc, { lineWidth: -1 });
fs.writeFileSync(openapiYamlPath, output, "utf8");

console.log("✅ contentType: text/plain を削除しました。");
