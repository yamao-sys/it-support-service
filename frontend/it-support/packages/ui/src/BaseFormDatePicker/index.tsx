import { memo } from "react";
import { Control, Controller, FieldValues, Path } from "react-hook-form";
import ReactDatePicker from "react-datepicker";

type Props<T extends FieldValues> = {
  id: string;
  label: string;
  control: Control<T>;
  name: Path<T>;
  validationErrors: string[];
};

const BaseFormDatePicker = memo(function BaseFormDatePicker<T extends FieldValues>({
  id,
  label,
  control,
  name,
  validationErrors,
}: Props<T>) {
  return (
    <>
      <label
        htmlFor={id}
        className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'
      >
        <span className='font-bold'>{label}</span>
      </label>
      <Controller
        control={control}
        name={name}
        render={({ field: { value, ...fieldProps } }) => (
          <ReactDatePicker
            id={id}
            {...fieldProps}
            dateFormat='yyyy-MM-dd'
            className='w-full px-3 py-2.5 focus-visible:outline-none bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
            selected={value}
          />
        )}
      />
      {!!validationErrors.length && (
        <div className='w-full pt-5 text-left'>
          {validationErrors.map((message, i) => (
            <p key={i} className='text-red-400'>
              {message}
            </p>
          ))}
        </div>
      )}
    </>
  );
});

export default BaseFormDatePicker;
