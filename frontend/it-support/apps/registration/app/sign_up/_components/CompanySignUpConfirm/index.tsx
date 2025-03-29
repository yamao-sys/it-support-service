"use client";

import { FC, useCallback } from "react";
import BaseImage from "@/components/BaseImage";
import { PhaseType } from "../../_types";
import { useCompanySignUpContext } from "../../_contexts/useCompanySignUpContext";
import BaseButton from "@repo/ui/BaseButton";
import { postCompanySignUp } from "@/apis/companies.api";

type Props = {
  togglePhase: (newPhase: PhaseType) => void;
};

const CompanySignUpConfirm: FC<Props> = ({ togglePhase }: Props) => {
  const { companySignUpInputs } = useCompanySignUpContext();

  const handleBackToInput = () => togglePhase("input");

  const handleSignUp = useCallback(async () => {
    const response = await postCompanySignUp(companySignUpInputs);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (Object.keys(response.errors).length === 0) {
      togglePhase("thanks");
      return;
    }

    throw Error("invalid company sign up input");
  }, [companySignUpInputs, togglePhase]);

  return (
    <>
      <h3 className='w-full text-center text-2xl font-bold'>企業登録の入力内容</h3>

      <div className='flex w-full justify-around mt-16'>
        <div className='w-1/2 align-middle'>企業名: </div>
        <div className='w-1/2 align-middle'>{companySignUpInputs.name}</div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>メールアドレス: </div>
        <div className='w-1/2 align-middle'>{companySignUpInputs.email}</div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>パスワード: </div>
        <div className='w-1/2 align-middle'>{"*".repeat(companySignUpInputs.password.length)}</div>
      </div>
      <div className='flex w-full justify-around mt-8'>
        <div className='w-1/2 align-middle'>確定申告書(コピー): </div>
        <div className='w-1/2 align-middle'>
          <BaseImage file={companySignUpInputs.finalTaxReturn} />
        </div>
      </div>

      <div className='flex w-full justify-around mt-16'>
        <BaseButton
          borderColor='border-gray-500'
          bgColor='bg-gray-500'
          label='入力へ戻る'
          onClick={handleBackToInput}
        />
        <BaseButton
          borderColor='border-green-500'
          bgColor='bg-green-500'
          label='登録する'
          onClick={handleSignUp}
        />
      </div>
    </>
  );
};

export default CompanySignUpConfirm;
