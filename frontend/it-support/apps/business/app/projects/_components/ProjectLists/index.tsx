"use client";

import { Project } from "@/apis";
import { getProjects } from "@/services/project";
import Link from "next/link";
import { FC, useCallback, useEffect, useRef, useState } from "react";

type Props = {
  initialProjects: Project[];
  initialNextPageToken: number;
};

const ProjectLists: FC<Props> = ({ initialProjects, initialNextPageToken }: Props) => {
  const [displayProjects, setDisplayProjects] = useState<Project[]>(initialProjects);
  const [nextPageToken, setNextPageToken] = useState(initialNextPageToken);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const loader = useRef<HTMLDivElement | null>(null);

  // NOTE: 無限スクロールによるデータ取得
  const fetchItems = useCallback(async () => {
    setLoading(true);

    const res = await getProjects(String(nextPageToken));
    if (res === undefined) {
      setLoading(false);
      throw new Error("Failed to fetch projects");
    }
    setDisplayProjects((prev) => [...prev, ...res.projects]);

    // 最後まで取得したか確認
    if (Number(res.nextPageToken) === 0) {
      setHasMore(false);
    }

    setNextPageToken(Number(res.nextPageToken));
    setLoading(false);
  }, [nextPageToken]);

  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        const entry = entries[0];
        if (entry?.isIntersecting && !loading && hasMore) {
          fetchItems();
        }
      },
      { threshold: 1.0 },
    );

    const currentLoader = loader.current;
    if (currentLoader) {
      observer.observe(currentLoader);
    }

    return () => {
      if (currentLoader) observer.unobserve(currentLoader);
    };
  }, [loading, hasMore, fetchItems]);

  return (
    <>
      <div className='w-full mt-4 md:mt-16'>
        {displayProjects.length === 0 ? (
          <p className='text-center'>まだ案件が未登録です</p>
        ) : (
          displayProjects.map((project) => (
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

        {/* ローディング表示 */}
        {loading && <div style={{ textAlign: "center", padding: "10px" }}>Loading...</div>}

        {/* 読み込み終了メッセージ */}
        {!hasMore && (
          <div style={{ textAlign: "center", padding: "10px", color: "gray" }}>
            これ以上データはありません。
          </div>
        )}

        {/* 監視対象 */}
        <div ref={loader} style={{ height: "100px" }} />
      </div>
    </>
  );
};

export default ProjectLists;
