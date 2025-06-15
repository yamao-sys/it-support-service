"use client";

import { ToProject } from "@/apis";
import { ListRow } from "@/components/ui/ListRow";
import { getToProjects } from "@/services/toProject";
import Link from "next/link";
import { FC, useCallback, useEffect, useRef, useState } from "react";
import ProposalStatusBadge from "../ProposalStatusBadge";

type Props = {
  initialProjects: ToProject[];
  initialNextPageToken: number;
};

const ToProjectLists: FC<Props> = ({ initialProjects, initialNextPageToken }: Props) => {
  const [displayProjects, setDisplayProjects] = useState<ToProject[]>(initialProjects);
  const [nextPageToken, setNextPageToken] = useState(initialNextPageToken);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const loader = useRef<HTMLDivElement | null>(null);

  // NOTE: 無限スクロールによるデータ取得
  const fetchItems = useCallback(async () => {
    setLoading(true);

    const res = await getToProjects(String(nextPageToken));
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
            <ListRow key={project.id}>
              <div className='flex flex-col gap-1'>
                <div className='text-lg font-semibold text-gray-800 break-words'>
                  {project.title}
                </div>
                <ProposalStatusBadge status={project.proposalStatus} />
              </div>
              <Link
                className='py-2 px-8 border-green-500 bg-green-500 rounded-xl text-white hover:bg-green-700 transition'
                href={`/to_projects/${project.id}`}
                target='_blank'
              >
                詳細を見る
              </Link>
            </ListRow>
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

export default ToProjectLists;
