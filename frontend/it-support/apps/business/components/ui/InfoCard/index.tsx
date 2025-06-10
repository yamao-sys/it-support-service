import { FC } from "react";

type Props = {
  title: string;
  value: string;
};

export const InfoCard: FC<Props> = ({ title, value }: Props) => {
  return (
    <div className='bg-gray-50 p-4 rounded-xl shadow-sm border'>
      <div className='text-xs text-gray-500'>{title}</div>
      <div className='text-sm font-medium text-gray-900 mt-1'>{value}</div>
    </div>
  );
};
