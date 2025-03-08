"use client";

import { FC, useCallback } from "react";
import { Phase } from "../SupporterSignUpForm";
import { useSupporterSignUpContext } from "../../_contexts/useSupporterSignUpContext";
import BaseImage from "@/components/BaseImage";
import { postSignUp } from "../../_actions/supporters";

type Props = {
  togglePhase: (newPhase: Phase) => void;
};

const SupporterSignUpConfirm: FC<Props> = ({ togglePhase }: Props) => {
  const { supporterSignUpInputs } = useSupporterSignUpContext();

  const handleBackToInput = () => togglePhase("input");

  const handleSignUp = useCallback(async () => {
    const response = await postSignUp(supporterSignUpInputs);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (Object.keys(response.errors).length === 0) {
      togglePhase("thanks");
      return;
    }

    throw Error("invalid supporter sign up input");
  }, [supporterSignUpInputs, togglePhase]);

  return (
    <>
      <h3 className='w-full text-center text-2xl font-bold'>サポータ登録入力内容</h3>

      <div className='flex w-full justify-around mt-16'>
        <div className='w-1/2 align-middle'>ユーザ名: </div>
        <div className='w-1/2 align-middle'>{`${supporterSignUpInputs.lastName} ${supporterSignUpInputs.firstName}`}</div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>メールアドレス: </div>
        <div className='w-1/2 align-middle'>{supporterSignUpInputs.email}</div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>パスワード: </div>
        <div className='w-1/2 align-middle'>
          {"*".repeat(supporterSignUpInputs.password.length)}
        </div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>身分証明書(表): </div>
        <div className='w-1/2 align-middle'>
          <BaseImage file={supporterSignUpInputs.frontIdentification} />
        </div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>身分証明書(裏): </div>
        <div className='w-1/2 align-middle'>
          <BaseImage file={supporterSignUpInputs.backIdentification} />
        </div>
      </div>

      <div className='flex w-full justify-around mt-16'>
        <button
          className='py-2 px-8 border-gray-500 bg-gray-500 rounded-xl text-white'
          onClick={handleBackToInput}
        >
          入力へ戻る
        </button>
        <button
          className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white'
          onClick={handleSignUp}
        >
          登録する
        </button>
      </div>
    </>
  );
};

export default SupporterSignUpConfirm;
