import { useQuery } from "react-query";
import { useParams } from "react-router-dom";
import { Post } from "./Post";

export interface ApiResponse<T> {
  success: string;
  data: T;
}

export const Posts = (props: any) => {
  const { slug } = useParams();
  const { isLoading, data } = useQuery<ApiResponse<Post[]>>(
    "posts",
    async () => {
      const response = await fetch("http://localhost:8080/posts");
      return response.json();
    },
    { refetchOnWindowFocus: false }
  );

  if (isLoading) {
    return <></>;
  }

  return (
    <div>
      {data?.data
        .sort((a, b) => (a.createdAt < b.createdAt ? 1 : -1))
        .map((post) => {
          return <Post key={post.slug} post={post} isPreview></Post>;
        })}
    </div>
  );
};
