import { FC, useEffect, useState } from "react";

type Props = {
  file: File | Blob | undefined;
};

const BaseImage: FC<Props> = ({ file }: Props) => {
  const [imageSource, setImageSource] = useState("");

  const generateImageSource = (binary: File | Blob) => {
    const fileReader = new FileReader();
    fileReader.onload = () => {
      setImageSource(fileReader.result as string);
    };
    fileReader.readAsDataURL(binary);
  };

  useEffect(() => {
    if (!file) return;

    generateImageSource(file);
  }, [file]);

  return (
    <>
      <div
        className='mb-2'
        style={{
          border: "black 3px dotted",
          display: "flex",
          borderRadius: 12,
          aspectRatio: "4 / 3",
          justifyContent: "center",
          alignItems: "center",
          overflow: "hidden",
          cursor: "pointer",
        }}
      >
        {/* 画像があればプレビューし、なければ「+ 画像をアップロード」を表示 */}
        {/* eslint-disable @next/next/no-img-element */}
        {imageSource ? (
          <img
            src={imageSource}
            alt='アップロード画像'
            style={{ objectFit: "contain", width: "100%", height: "100%" }}
          />
        ) : (
          "-"
        )}
        {/* eslint-enable @next/next/no-img-element */}
      </div>
    </>
  );
};

export default BaseImage;
