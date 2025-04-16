import { Project } from "@/types";
import Link from "next/link";
import { FC } from "react";

type Props = {
  projects: Project[];
};

const ProjectLists: FC<Props> = ({ projects }: Props) => {
  return (
    <>
      <div className='w-full mt-4 md:mt-16'>
        {projects.length === 0 ? (
          <p className='text-center'>まだ案件が未登録です</p>
        ) : (
          projects.map((project) => (
            <div
              key={project.id}
              className='bg-white shadow-md rounded-2xl p-4 mb-4 flex items-center justify-between'
            >
              <div className='flex flex-col gap-1'>
                <div className='text-lg font-semibold text-gray-800 break-words'>
                  {project.title}
                </div>
                <div className='text-sm text-gray-600'>
                  <span className='text-green-600 font-medium'>
                    {project.isActive ? "公開中" : "未公開"}
                  </span>
                </div>
                <div className='text-sm text-gray-600'>
                  応募数：<span className='font-semibold'>12件</span>
                </div>
              </div>
              <Link
                className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white hover:bg-green-700 transition'
                href={`/projects/${project.id}`}
                target='_blank'
              >
                編集
              </Link>
            </div>
          ))
        )}
      </div>
    </>
  );
};

export default ProjectLists;
