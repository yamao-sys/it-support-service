import ToProjectListsContainer from "@/features/to-projects/Container/ToProjectListsContainer";
import { Suspense } from "react";

export default async function ToProjectListPage() {
  return (
    <>
      <Suspense fallback={<p>loading...</p>}>
        <ToProjectListsContainer />
      </Suspense>
    </>
  );
}
