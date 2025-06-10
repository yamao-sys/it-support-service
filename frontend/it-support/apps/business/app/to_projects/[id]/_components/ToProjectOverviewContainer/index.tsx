import { getToProject } from "@/services/toProject";
import { FC } from "react";
import ToProjectOverview from "../ToProjectOverview";

type Props = {
  id: number;
};

const ToProjectOverviewContainer: FC<Props> = async ({ id }: Props) => {
  const res = await getToProject(id);
  if (res === undefined) {
    throw new Error("Failed to fetch projects");
  }

  return (
    <>
      <ToProjectOverview project={res.project} />
    </>
  );
};

export default ToProjectOverviewContainer;
