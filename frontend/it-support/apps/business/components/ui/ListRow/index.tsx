import { FC } from "react";

type Props = {
  children: React.ReactNode;
};

export const ListRow: FC<Props> = ({ children }: Props) => {
  return (
    <div className='bg-white shadow-md rounded-2xl p-4 mb-4 flex items-center justify-between'>
      {children}
    </div>
  );
};
