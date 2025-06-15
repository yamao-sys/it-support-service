import { Suspense } from "react";

type ToProjectPlanProposalPageProps = {
  params: Promise<{
    id: string;
  }>;
};

export default async function ToProjectPlanProposalPage({
  params,
}: ToProjectPlanProposalPageProps) {
  const { id } = await params;

  return (
    <>
      <Suspense fallback={<>loading...</>}>
        {/* <ToProjectOverviewContainer id={Number(id)} /> */}
      </Suspense>
    </>
  );
}
