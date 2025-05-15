import ProjectEditTemplate from "./_components/ProjectEditTemplate";
import { Suspense } from "react";

type ProjectEditPageProps = {
  params: Promise<{
    id: string;
  }>;
};

export default async function ProjectEditPage({ params }: ProjectEditPageProps) {
  const { id } = await params;

  return (
    <>
      <Suspense fallback={<>loading...</>}>
        <ProjectEditTemplate projectId={Number(id)} />
      </Suspense>
    </>
  );
}
