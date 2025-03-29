"use client";

import { FC, useCallback, useState } from "react";
import { FormType, PhaseType, SupporterSignUpValidationError } from "../../_types";
import FormTypeSelector from "../FormTypeSelector";
import {
  useSupporterSignUpContext,
  useSupporterSignUpSetContext,
} from "../../_contexts/useSupporterSignUpContext";
import BaseImageInputForm from "@/components/BaseImageInputForm";
import BaseFormInput from "@repo/ui/BaseFormInput";
import BaseButton from "@repo/ui/BaseButton";
import { postSupporterValidateSignUp } from "@/apis/supporters.api";

type Props = {
  formType: FormType;
  togglePhase: (newPhase: PhaseType) => void;
  switchFormType: (newFormType: FormType) => void;
};

const INITIAL_VALIDATION_ERRORS = {
  firstName: [],
  lastName: [],
  email: [],
  password: [],
  birthday: [],
  frontIdentification: [],
  backIdentification: [],
};

const SupporterSignUpInput: FC<Props> = ({ formType, togglePhase, switchFormType }: Props) => {
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

    const response = await postSupporterValidateSignUp(supporterSignUpInputs);

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
      <FormTypeSelector formType={formType} switchFormType={switchFormType} />

      <h3 className='mt-16 w-full text-center text-2xl font-bold'>サポータ登録フォーム</h3>

      <div className='flex justify-between'>
        <div className='mt-8' style={{ width: "45%" }}>
          <BaseFormInput
            id='last-name'
            label='姓'
            name='lastName'
            type='text'
            value={supporterSignUpInputs.lastName}
            onChange={setSupporterSignUpTextInput}
            validationErrorMessages={validationErrors.lastName ?? []}
          />
        </div>

        <div className='mt-8' style={{ width: "45%" }}>
          <BaseFormInput
            id='first-name'
            label='名'
            name='firstName'
            type='text'
            value={supporterSignUpInputs.firstName}
            onChange={setSupporterSignUpTextInput}
            validationErrorMessages={validationErrors.firstName ?? []}
          />
        </div>
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='email'
          label='Email'
          name='email'
          type='email'
          value={supporterSignUpInputs.email}
          onChange={setSupporterSignUpTextInput}
          validationErrorMessages={validationErrors.email ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='password'
          label='パスワード'
          name='password'
          type='password'
          value={supporterSignUpInputs.password}
          onChange={setSupporterSignUpTextInput}
          validationErrorMessages={validationErrors.password ?? []}
        />
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

export default SupporterSignUpInput;
