import { JSX, memo } from "react";

type Props = {
  label: string;
  id: string;
  validationErrorMessages: string[];
} & JSX.IntrinsicElements["textarea"];

const BaseFormTextarea = memo(function BaseFormTextarea({
  label,
  id,
  validationErrorMessages,
  ...props
}: Props) {
  return (
    <>
      <label
        htmlFor={id}
        className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
      >
        <span className='font-bold'>{label}</span>
      </label>
      <textarea
        id={id}
        rows={16}
        className='block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 resize-none'
        {...props}
      />
      {validationErrorMessages.length > 0 && (
        <div className='w-full pt-5 text-left'>
          {validationErrorMessages.map((message, i) => (
            <p key={i} className='text-red-400'>
              {message}
            </p>
          ))}
        </div>
      )}
    </>
  );
});

export default BaseFormTextarea;
