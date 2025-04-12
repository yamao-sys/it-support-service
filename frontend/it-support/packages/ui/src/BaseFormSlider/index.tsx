import { memo } from "react";
import { FieldValues, Path, Control, Controller } from "react-hook-form";
import BaseSlider from "../BaseSlider";

type Props<T extends FieldValues> = {
  label: string;
  control: Control<T>;
  name: Path<T>;
  validationErrors: string[];
};

const BaseFormSlider = memo(function BaseFormSlider<T extends FieldValues>({
  label,
  control,
  name,
  validationErrors,
}: Props<T>) {
  return (
    <>
      <Controller
        control={control}
        name={name}
        render={({ field }) => <BaseSlider label={label} {...field} />}
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

export default BaseFormSlider;
