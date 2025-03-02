import openapiTS, { astToString } from "openapi-typescript";
import fs from "fs";
import ts from "typescript";

const schemaPath = process.argv[2];
const outPutPath = process.argv[3];

// NOTE: OpenAPI スキーマをロード
const mySchema = fs.readFileSync(schemaPath, "utf-8");

const BLOB = ts.factory.createTypeReferenceNode(ts.factory.createIdentifier("Blob"));
const NULL = ts.factory.createLiteralTypeNode(ts.factory.createNull());

// NOTE: 型生成オプションを設定
/* eslint-disable @typescript-eslint/no-unused-vars */
const typeFileContent = await openapiTS(mySchema, {
  transform(schemaObject, metadata) {
    if ("format" in schemaObject && schemaObject.format === "binary") {
      return schemaObject.nullable ? ts.factory.createUnionTypeNode([BLOB, NULL]) : BLOB;
    }
    return undefined; // NOTE: 他の型に関してはデフォルト処理
  },
});
/* eslint-enable @typescript-eslint/no-unused-vars */

const contents = astToString(typeFileContent);

// （任意）ファイルに書き込み
// await fs.promises.mkdir(outPutPath);
fs.writeFileSync(`${outPutPath}apiSchema.ts`, contents);
