import BaseContainer from "@/components/BaseContailer";
import { FC, ReactNode } from "react";
import { Phase } from "../SupporterSignUpForm";

type Props = {
  phase: Phase;
  children: ReactNode;
};

const SignUpLayout: FC<Props> = ({ phase, children }: Props) => {
  return (
    <div className='p-4 md:p-16'>
      <BaseContainer containerWidth='w-4/5 md:w-3/5'>
        <div className='flex justify-between mb-16'>
          <div>
            <span className={phase === "input" ? "text-blue-300" : "text-gray-300"}>
              登録情報の入力
            </span>
          </div>
          <div className='text-gray-300'>&gt;&gt;</div>

          <div>
            <span className={phase === "confirm" ? "text-blue-300" : "text-gray-300"}>
              登録情報の確認
            </span>
          </div>
          <div className='text-gray-300'>&gt;&gt;</div>

          <div>
            <span className={phase === "thanks" ? "text-blue-300" : "text-gray-300"}>登録完了</span>
          </div>
        </div>

        {children}
      </BaseContainer>
    </div>
  );
};

export default SignUpLayout;
