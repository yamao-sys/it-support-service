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
        <label
          htmlFor='company-name'
          className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
        >
          <span className='font-bold'>企業名</span>
        </label>
        <input
          id='company-name'
          name='companyName'
          type='text'
          className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
          value={companySignUpInputs.name}
          onChange={(e) => updateSignUpInput({ name: e.target.value })}
        />
        {validationErrors.name && (
          <div className='w-full pt-5 text-left'>
            {validationErrors.name.map((message, i) => (
              <p key={i} className='text-red-400'>
                {message}
              </p>
            ))}
          </div>
        )}
      </div>

      <div className='mt-8'>
        <label
          htmlFor='company-email'
          className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
        >
          <span className='font-bold'>Email</span>
        </label>
        <input
          id='company-email'
          name='companyEmail'
          type='email'
          className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
          value={companySignUpInputs.email}
          onChange={(e) => updateSignUpInput({ email: e.target.value })}
        />
        {validationErrors.email && (
          <div className='w-full pt-5 text-left'>
            {validationErrors.email.map((message, i) => (
              <p key={i} className='text-red-400'>
                {message}
              </p>
            ))}
          </div>
        )}
      </div>

      <div className='mt-8'>
        <label
          htmlFor='company-password'
          className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
        >
          <span className='font-bold'>パスワード</span>
        </label>
        <input
          id='company-password'
          name='Companypassword'
          type='password'
          className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
          value={companySignUpInputs.password}
          onChange={(e) => updateSignUpInput({ password: e.target.value })}
        />
        {validationErrors.password && (
          <div className='w-full pt-5 text-left'>
            {validationErrors.password.map((message, i) => (
              <p key={i} className='text-red-400'>
                {message}
              </p>
            ))}
          </div>
        )}
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
          <button
            type='button'
            className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white'
            onClick={handleValidateSignUp}
          >
            確認画面へ
          </button>
        </div>
      </div>
    </>
  );
};

export default CompanySignUpInput;
