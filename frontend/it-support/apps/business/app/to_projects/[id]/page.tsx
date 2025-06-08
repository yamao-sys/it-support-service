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
        <>{id}</>
      </Suspense>
    </>
  );
}
