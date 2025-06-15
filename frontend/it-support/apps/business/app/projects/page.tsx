import ProjectListsContainer from "@/features/projects/components/ProjectListsContainer";
import { Suspense } from "react";

export default async function ProjectListPage() {
  return (
    <>
      <Suspense fallback={<p>loading...</p>}>
        <ProjectListsContainer />
      </Suspense>
    </>
  );
}
