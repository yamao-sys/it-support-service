import { Suspense } from "react";
import ProjectListsContainer from "./_components/ProjectListsContainer";

export default async function ProjectListPage() {
  return (
    <>
      <Suspense fallback={<p>loading...</p>}>
        <ProjectListsContainer />
      </Suspense>
    </>
  );
}
