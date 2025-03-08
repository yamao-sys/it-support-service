"use client";

import { FC, useCallback, useState } from "react";
import { postValidateSignUp } from "../../_actions/supporters";
import { components } from "@/apis/generated/supporters/apiSchema";
import {
  useSupporterSignUpContext,
  useSupporterSignUpSetContext,
} from "../../_contexts/useSupporterSignUpContext";
import { Phase } from "../SupporterSignUpForm";
import BaseImageInputForm from "@/components/BaseImageInputForm";

export type SupporterSignUpInput =
  components["requestBodies"]["SignUpInput"]["content"]["multipart/form-data"];

type SupporterSignUpValidationError =
  components["responses"]["SignUpResponse"]["content"]["application/json"]["errors"];

const INITIAL_VALIDATION_ERRORS = {
  firstName: [],
  lastName: [],
  email: [],
  password: [],
  birthday: [],
  frontIdentification: [],
  backIdentification: [],
};

type Props = {
  togglePhase: (newPhase: Phase) => void;
};

const SupporterRegistrationForm: FC<Props> = ({ togglePhase }: Props) => {
  const { supporterSignUpInputs } = useSupporterSignUpContext();
  const { updateSignUpInput, clearIdentificationKey } = useSupporterSignUpSetContext();

  const [validationErrors, setValidationErrors] =
    useState<SupporterSignUpValidationError>(INITIAL_VALIDATION_ERRORS);

  const setSupporterSignUpTextInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      updateSignUpInput({ [e.target.name]: e.target.value });
    },
    [updateSignUpInput],
  );

  const setSupporterSignUpFileInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>, file: File) => {
      updateSignUpInput({ [e.target.name]: file });
    },
    [updateSignUpInput],
  );

  const clearSupporterSignUpFileInput = useCallback(
    (key: string) => {
      if (key === "frontIdentification" || key === "backIdentification") {
        clearIdentificationKey(key);
      }
    },
    [clearIdentificationKey],
  );

  const handleValidateSignUp = useCallback(async () => {
    setValidationErrors(INITIAL_VALIDATION_ERRORS);

    const response = await postValidateSignUp(supporterSignUpInputs);

    // バリデーションエラーがなければ、確認画面へ遷移
    if (Object.keys(response.errors).length === 0) {
      togglePhase("confirm");
      return;
    }

    // NOTE: バリデーションエラーの格納と入力パスワードのリセット
    setValidationErrors(response.errors);
    updateSignUpInput({ password: "" });
  }, [setValidationErrors, supporterSignUpInputs, togglePhase, updateSignUpInput]);

  return (
    <>
      <div className='flex justify-between'>
        <div className='mt-8' style={{ width: "45%" }}>
          <label
            htmlFor='last-name'
            className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
          >
            <span className='font-bold'>姓</span>
          </label>
          <input
            id='last-name'
            name='lastName'
            type='text'
            className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
            value={supporterSignUpInputs.lastName}
            onChange={setSupporterSignUpTextInput}
          />
          {validationErrors.lastName && (
            <div className='w-full pt-5 text-left'>
              {validationErrors.lastName.map((message, i) => (
                <p key={i} className='text-red-400'>
                  {message}
                </p>
              ))}
            </div>
          )}
        </div>

        <div className='mt-8' style={{ width: "45%" }}>
          <label
            htmlFor='first-name'
            className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
          >
            <span className='font-bold'>名</span>
          </label>
          <input
            id='first-name'
            name='firstName'
            type='text'
            className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
            value={supporterSignUpInputs.firstName}
            onChange={setSupporterSignUpTextInput}
          />
          {validationErrors.firstName && (
            <div className='w-full pt-5 text-left'>
              {validationErrors.firstName.map((message, i) => (
                <p key={i} className='text-red-400'>
                  {message}
                </p>
              ))}
            </div>
          )}
        </div>
      </div>

      <div className='mt-8'>
        <label
          htmlFor='email'
          className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
        >
          <span className='font-bold'>Email</span>
        </label>
        <input
          id='email'
          name='email'
          type='email'
          className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
          value={supporterSignUpInputs.email}
          onChange={setSupporterSignUpTextInput}
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
          htmlFor='password'
          className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
        >
          <span className='font-bold'>パスワード</span>
        </label>
        <input
          id='password'
          name='password'
          type='password'
          className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
          value={supporterSignUpInputs.password}
          onChange={setSupporterSignUpTextInput}
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
          id='front-identification'
          name='frontIdentification'
          label='身分証明書(表)'
          initialFileInput={supporterSignUpInputs.frontIdentification}
          onChange={setSupporterSignUpFileInput}
          onCancel={clearSupporterSignUpFileInput}
          validationErrorMessages={validationErrors.frontIdentification ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseImageInputForm
          id='back-identification'
          name='backIdentification'
          label='身分証明書(裏)'
          initialFileInput={supporterSignUpInputs.backIdentification}
          onChange={setSupporterSignUpFileInput}
          onCancel={clearSupporterSignUpFileInput}
          validationErrorMessages={validationErrors.backIdentification ?? []}
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

export default SupporterRegistrationForm;
