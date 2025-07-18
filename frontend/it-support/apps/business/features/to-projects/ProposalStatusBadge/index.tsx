"use client";

import { ToProjectProposalStatusEnum } from "@/apis";
import { FC, memo, useMemo } from "react";

type Props = {
  status: ToProjectProposalStatusEnum;
};

const ProposalStatusBadge: FC<Props> = memo(function ProposalStatusBadge({ status }: Props) {
  const [statusBadgeColor, statusBadgeText] = useMemo(() => {
    switch (status) {
      case "NOT PROPOSED":
        return ["bg-gray-100 text-gray-800", "未提案"];
      case "TEMPORARY CREATING":
        return ["bg-yellow-100 text-yellow-800", "一時作成中"];
      case "PROPOSED":
        return ["bg-blue-100 text-blue-800", "提案済み"];
    }
  }, [status]);

  return (
    <span className={`inline-block text-xs font-medium px-2 py-1 rounded ${statusBadgeColor}`}>
      {statusBadgeText}
    </span>
  );
});

export default ProposalStatusBadge;
