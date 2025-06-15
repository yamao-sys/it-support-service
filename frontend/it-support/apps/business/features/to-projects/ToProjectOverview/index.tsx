import { ToProject } from "@/apis";
import { InfoCard } from "@/components/ui/InfoCard";
import { getDateString } from "@/lib/date";
import { FC } from "react";

type Props = {
  project: ToProject;
};

const ToProjectOverview: FC<Props> = ({ project }: Props) => {
  return (
    <>
      <div className='bg-white shadow p-6 rounded-2xl'>
        <h1 className='text-2xl font-semibold text-gray-900'>{project.title}</h1>
      </div>

      <div className='mt-4'>
        <h2 className='text-lg font-medium text-gray-800'>案件概要</h2>
        <p className='text-gray-700 mt-2 whitespace-pre-line'>{project.description}</p>
      </div>

      <div className='mt-6 grid grid-cols-1 md:grid-cols-2 gap-4'>
        <InfoCard title='開始日' value={getDateString(project.startDate)} />
        <InfoCard title='終了日' value={getDateString(project.endDate)} />
        <InfoCard
          title='予算（下限）'
          value={project.minBudget ? `¥${project.minBudget.toLocaleString("")}` : "未定"}
        />
        <InfoCard
          title='予算（上限）'
          value={project.maxBudget ? `¥${project.maxBudget.toLocaleString()}` : "未定"}
        />
      </div>
    </>
  );
};

export default ToProjectOverview;
