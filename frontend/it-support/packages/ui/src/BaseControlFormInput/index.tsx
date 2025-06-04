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

function BaseControlFormInputInner<T extends FieldValues>({
  id,
  type = "text",
  label,
  control,
  name,
  validationErrors,
}: Props<T>) {
  return (
    <Controller
      control={control}
      name={name}
      render={({ field }) => {
        const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
          const value = e.target.value;
          const parsedValue =
            type === "number" ? (value === "" ? undefined : Number(value)) : value;
          field.onChange(parsedValue);
        };

        return (
          <BaseFormInput
            id={id}
            label={label}
            type={type}
            {...field}
            value={field.value ?? ""}
            onChange={handleChange}
            validationErrorMessages={validationErrors}
          />
        );
      }}
    />
  );
}

const BaseControlFormInput = memo(BaseControlFormInputInner) as typeof BaseControlFormInputInner;

export default BaseControlFormInput;
