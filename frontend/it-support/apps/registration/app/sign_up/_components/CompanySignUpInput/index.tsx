"use client";

import { FC, useCallback, useState } from "react";
import { FormType, PhaseType, CompanySignUpValidationError } from "../../_types";
import FormTypeSelector from "../FormTypeSelector";
import { postValidateSignUp } from "../../_actions/companies";
import BaseImageInputForm from "@/components/BaseImageInputForm";
import {
  useCompanySignUpContext,
  useCompanySignUpSetContext,
} from "../../_contexts/useCompanySignUpContext";
import BaseFormInput from "@/components/BaseFormInput";
import BaseButton from "@/components/BaseButton";

type Props = {
  formType: FormType;
  togglePhase: (newPhase: PhaseType) => void;
  switchFormType: (newFormType: FormType) => void;
};

const INITIAL_VALIDATION_ERRORS = {
  name: [],
  email: [],
  password: [],
  finalTaxReturn: [],
};

const CompanySignUpInput: FC<Props> = ({ formType, togglePhase, switchFormType }: Props) => {
  const { companySignUpInputs } = useCompanySignUpContext();
  const { updateSignUpInput, clearIdentificationKey } = useCompanySignUpSetContext();

  const [validationErrors, setValidationErrors] =
    useState<CompanySignUpValidationError>(INITIAL_VALIDATION_ERRORS);

  const setSignUpFileInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>, file: File) => {
      updateSignUpInput({ [e.target.name]: file });
    },
    [updateSignUpInput],
  );

  const clearSignUpFileInput = useCallback(
    (key: string) => {
      if (key === "finalTaxReturn") {
        clearIdentificationKey(key);
      }
    },
    [clearIdentificationKey],
  );

  const handleValidateSignUp = useCallback(async () => {
    setValidationErrors(INITIAL_VALIDATION_ERRORS);

    const response = await postValidateSignUp(companySignUpInputs);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (Object.keys(response.errors).length === 0) {
      togglePhase("confirm");
      return;
    }

    // NOTE: バリデーションエラーの格納と入力パスワードのリセット
    setValidationErrors(response.errors);
    updateSignUpInput({ password: "" });
  }, [setValidationErrors, companySignUpInputs, togglePhase, updateSignUpInput]);

  return (
    <>
      <FormTypeSelector formType={formType} switchFormType={switchFormType} />

      <h3 className='mt-16 w-full text-center text-2xl font-bold'>企業登録フォーム</h3>

      <div className='mt-8'>
        <BaseFormInput
          id='company-name'
          label='企業名'
          name='companyName'
          type='text'
          value={companySignUpInputs.name}
          onChange={(e) => updateSignUpInput({ name: e.target.value })}
          validationErrorMessages={validationErrors.name ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='company-email'
          label='Email'
          name='companyEmail'
          type='email'
          value={companySignUpInputs.email}
          onChange={(e) => updateSignUpInput({ email: e.target.value })}
          validationErrorMessages={validationErrors.email ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='company-password'
          label='パスワード'
          name='companyPassword'
          type='password'
          value={companySignUpInputs.password}
          onChange={(e) => updateSignUpInput({ password: e.target.value })}
          validationErrorMessages={validationErrors.password ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseImageInputForm
          id='final-tax-return'
          name='finalTaxReturn'
          label='確定申告書(コピー)'
          initialFileInput={companySignUpInputs.finalTaxReturn}
          onChange={setSignUpFileInput}
          onCancel={clearSignUpFileInput}
          validationErrorMessages={validationErrors.finalTaxReturn ?? []}
        />
      </div>

      <div className='w-full flex justify-center'>
        <div className='mt-16'>
          <BaseButton
            borderColor='border-green-500'
            bgColor='bg-green-500'
            label='確認画面へ'
            onClick={handleValidateSignUp}
          />
        </div>
      </div>
    </>
  );
};

export default CompanySignUpInput;
