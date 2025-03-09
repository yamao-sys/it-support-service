import { FC, useEffect, useRef, useState } from "react";
import BaseImageInput from "../BaseImageInput";

type Props = {
  id: string;
  name: string;
  label: string;
  initialFileInput?: Blob | undefined;
  onChange: (e: React.ChangeEvent<HTMLInputElement>, file: File) => void;
  onCancel: (key: string) => void;
  validationErrorMessages: string[];
};

const BaseImageInputForm: FC<Props> = ({
  id,
  name,
  label,
  initialFileInput = undefined,
  onChange,
  onCancel,
  validationErrorMessages = [],
}: Props) => {
  const fileInputRef = useRef<HTMLInputElement>(null);

  // 添付画像を状態管理
  const [imageSource, setImageSource] = useState("");
  const [imageFile, setImageFile] = useState<Blob | File | null>(initialFileInput ?? null);

  const selectFile = () => {
    if (!fileInputRef.current) return;
    // ローカルフォルダーから画像選択のダイアログが表示される。
    fileInputRef.current.click();
  };

  const generateImageSource = (file: File | Blob) => {
    const fileReader = new FileReader();
    fileReader.onload = () => {
      setImageSource(fileReader.result as string);
    };
    fileReader.readAsDataURL(file);
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (!files) return;

    const file = files[0];
    if (!file) return;

    setImageFile(file);
    onChange(e, file);
  };

  // キャンセルボタンを押した際の処理
  const handleClickCancelButton = () => {
    setImageFile(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = "";
      onCancel(fileInputRef.current.name);
    }
  };

  useEffect(() => {
    if (!imageFile) return;

    generateImageSource(imageFile);
  }, [imageFile]);

  return (
    <div className='w-full md:w-3/4'>
      <label className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'>
        <span className='font-bold'>{label}</span>
      </label>

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
        onClick={selectFile}
      >
        {/* 画像があればプレビューし、なければ「+ 画像をアップロード」を表示 */}
        {/* eslint-disable @next/next/no-img-element */}
        {imageFile && imageSource ? (
          <img
            src={imageSource}
            alt='アップロード画像'
            style={{ objectFit: "contain", width: "100%", height: "100%" }}
          />
        ) : (
          `+ ${label}を選択`
        )}
        {/* eslint-enable @next/next/no-img-element */}
        <BaseImageInput ref={fileInputRef} id={id} name={name} onChange={handleFileChange} />
      </div>

      <div className='w-full flex justify-center'>
        <button
          type='button'
          className='py-2 px-8 mx-auto border-gray-500 bg-gray-500 rounded-xl text-white'
          onClick={handleClickCancelButton}
        >
          × キャンセル
        </button>
      </div>

      {validationErrorMessages.length > 0 && (
        <div className='w-full pt-5 text-left'>
          {validationErrorMessages.map((message, i) => (
            <p key={i} className='text-red-400'>
              {message}
            </p>
          ))}
        </div>
      )}
    </div>
  );
};

export default BaseImageInputForm;
