"use client";

import { ToProjectProposalStatusEnum } from "@/apis";
import { FC, useCallback } from "react";

type Props = {
  status: ToProjectProposalStatusEnum;
};

const ProposalStatusBadge: FC<Props> = ({ status }: Props) => {
  const statusBadge = useCallback((status: ToProjectProposalStatusEnum) => {
    switch (status) {
      case "NOT PROPOSED":
        return ["bg-gray-100 text-gray-800", "未提案"];
      case "TEMPORARY CREATING":
        return ["bg-yellow-100 text-yellow-800", "一時作成中"];
      case "PROPOSED":
        return ["bg-blue-100 text-blue-800", "提案済み"];
    }
  }, []);

  return (
    <span
      className={`inline-block text-xs font-medium px-2 py-1 rounded ${statusBadge(status)[0]}`}
    >
      {statusBadge(status)[1]}
    </span>
  );
};

export default ProposalStatusBadge;
