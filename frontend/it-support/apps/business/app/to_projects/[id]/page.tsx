import { Suspense } from "react";
import ToProjectOverviewContainer from "./_components/ToProjectOverviewContainer";

type ToProjectDetailPageProps = {
  params: Promise<{
    id: string;
  }>;
};

export default async function ToProjectDetailPage({ params }: ToProjectDetailPageProps) {
  const { id } = await params;

  return (
    <>
      <Suspense fallback={<>loading...</>}>
        <ToProjectOverviewContainer id={Number(id)} />
      </Suspense>
    </>
  );
}
