import ToProjectOverviewContainer from "@/features/to-projects/Container/ToProjectOverviewContainer";
import { Suspense } from "react";

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
