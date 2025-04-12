import { memo } from "react";
import { Control, Controller, FieldValues, Path } from "react-hook-form";
import BaseFormInput from "../BaseFormInput";

type Props<T extends FieldValues> = {
  id: string;
  type?: string;
  label: string;
  control: Control<T>;
  name: Path<T>;
  validationErrors: string[];
};

const BaseControlFormInput = memo(function BaseControlFormInput<T extends FieldValues>({
  id,
  type = "text",
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
        render={({ field }) => (
          <BaseFormInput
            id={id}
            label={label}
            type={type}
            {...field}
            value={field.value ?? ""}
            validationErrorMessages={validationErrors}
          />
        )}
      />
    </>
  );
});

export default BaseControlFormInput;
