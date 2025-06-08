import { Suspense } from "react";
import ToProjectListsContainer from "./_components/ToProjectListsContainer";

export default async function ToProjectListPage() {
  return (
    <>
      <Suspense fallback={<p>loading...</p>}>
        <ToProjectListsContainer />
      </Suspense>
    </>
  );
}
