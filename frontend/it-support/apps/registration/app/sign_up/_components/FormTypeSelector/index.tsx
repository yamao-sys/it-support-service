"useClient";

import { FC, useCallback } from "react";
import { FormType } from "../../_types";

type Props = {
  formType: FormType;
  switchFormType: (newFormType: FormType) => void;
};

const FormTypeSelector: FC<Props> = ({ formType, switchFormType }: Props) => {
  const switchToSupporterForm = useCallback(() => switchFormType("supporter"), [switchFormType]);
  const switchToCompanyForm = useCallback(() => switchFormType("company"), [switchFormType]);

  return (
    <>
      <div className='flex w-full justify-around'>
        <div className='flex w-2/5 items-center ps-4 border border-gray-200 rounded dark:border-gray-700'>
          <input
            checked={formType === "supporter"}
            id='supporter-form'
            type='radio'
            value=''
            name='supporter-form'
            className='w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600'
            onChange={switchToSupporterForm}
          />
          <label
            htmlFor='supporter-form'
            className='w-full py-4 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300'
          >
            サポータ
          </label>
        </div>
        <div className='flex w-2/5 items-center ps-4 border border-gray-200 rounded dark:border-gray-700'>
          <input
            checked={formType === "company"}
            id='company-form'
            type='radio'
            value=''
            name='company-form'
            className='w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600'
            onChange={switchToCompanyForm}
          />
          <label
            htmlFor='company-form'
            className='w-full py-4 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300'
          >
            企業
          </label>
        </div>
      </div>
    </>
  );
};

export default FormTypeSelector;
