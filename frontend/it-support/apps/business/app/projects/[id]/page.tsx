import ProjectEditTemplate from "@/features/projects/components/Template/ProjectEditTemplate";
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
