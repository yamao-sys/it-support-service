import BaseContainer from "@repo/ui/BaseContainer";

export default async function SignInLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className='p-4 md:p-16'>
      <BaseContainer containerWidth='w-4/5 md:w-3/5'>{children}</BaseContainer>
    </div>
  );
}
